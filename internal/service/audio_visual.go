package service

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/zkep/mygeektime/internal/global"
	"github.com/zkep/mygeektime/internal/model"
	"github.com/zkep/mygeektime/internal/types/geek"
	"github.com/zkep/mygeektime/lib/zhttp"
	"go.uber.org/zap"
)

var (
	SpecialCharacters = map[string]string{
		"|": "_",
		"?": "_",
		"？": "_",
		"：": "_",
		"/": "_",
		"*": "_",
		" ": "",
	}
)

func VerifyFileName(name string) string {
	for src, dest := range SpecialCharacters {
		name = strings.ReplaceAll(name, src, dest)
	}
	return name
}

type Part struct {
	Src  string
	Dest string
}

func Download(ctx context.Context, x *model.Task) error {
	var data geek.ArticleData
	if err := json.Unmarshal(x.Raw, &data); err != nil {
		return err
	}
	if data.Info.IsVideo {
		sort.Slice(data.Info.Video.HlsMedias, func(i, j int) bool {
			return data.Info.Video.HlsMedias[i].Size > data.Info.Video.HlsMedias[j].Size
		})
		hlsURL := data.Info.Video.HlsMedias[0].URL
		fileName := VerifyFileName(data.Info.Title)
		dir := path.Join(VerifyFileName(data.Product.Title), fileName)
		if source, err := Video(ctx, hlsURL, dir, fileName); err != nil {
			global.LOG.Error("download video", zap.Error(err), zap.String("taskId", x.TaskId))
			return err
		} else {
			x.Message = []byte(source)
		}
	} else if data.Info.Audio.DownloadURL != "" {
		if source, err := Audio(ctx, data.Info.Audio.DownloadURL,
			VerifyFileName(data.Product.Title), VerifyFileName(data.Info.Title)); err != nil {
			global.LOG.Error("download audio", zap.Error(err), zap.String("taskId", x.TaskId))
			return err
		} else {
			x.Message = []byte(source)
		}
	}
	return nil
}

func Video(ctx context.Context, hlsURL, dir, fileName string) (string, error) {
	retryCtx, retryCancel := context.WithTimeout(ctx, time.Minute*30)
	defer retryCancel()

	var hlsRaw []byte
	err := zhttp.R.
		Client(global.HttpClient).
		Before(func(r *http.Request) {
			r.Header.Set("Accept", "application/json, text/plain, */*")
			r.Header.Set("Content-Type", "application/json")
			r.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
			r.Header.Set("Cookie", global.GeekCookies)
			r.Header.Set("User-Agent", zhttp.RandomUserAgent())
			r.Header.Set("Referer", r.URL.String())
			r.Header.Set("Origin", "https://time.geekbang.org")
		}).
		After(func(r *http.Response) error {
			if zhttp.IsHTTPSuccessStatus(r.StatusCode) {
				raw, err := io.ReadAll(r.Body)
				if err != nil {
					return err
				}
				hlsRaw = raw
				return nil
			}
			if zhttp.IsHTTPStatusSleep(r.StatusCode) {
				time.Sleep(time.Second * 10)
			}
			if zhttp.IsHTTPStatusRetryable(r.StatusCode) {
				return fmt.Errorf("http status: %s", r.Status)
			}
			return zhttp.BreakRetryError(fmt.Errorf("http status: %s", r.Status))
		}).
		DoWithRetry(retryCtx, http.MethodGet, hlsURL, nil)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	bio := bufio.NewReader(bytes.NewReader(hlsRaw))
	ts := make([]Part, 0, 10)
	for {
		line, _, err1 := bio.ReadLine()
		if err1 != nil {
			break
		}
		l := string(line)
		srcURL, destName := "", ""
		if strings.HasPrefix(l, "#EXT-X-KEY:") {
			sps := strings.Split(l, `"`)
			srcURL = sps[1]
			destName = path.Join(dir, "key.key")
			l = fmt.Sprintf(`%s"file:///%s/key.key"`, sps[0], global.Storage.GetKey(dir))
		} else if strings.HasSuffix(l, ".ts") {
			srcURL = hlsURL[:strings.LastIndex(hlsURL, "/")+1] + l
			destName = path.Join(dir, l)
		}
		buff.WriteString(l + "\n")
		if len(srcURL) == 0 || len(destName) == 0 {
			continue
		}
		ts = append(ts, Part{srcURL, destName})
	}

	for _, t := range ts {
		dowloadURL, destName := t.Src, t.Dest
		err = zhttp.R.
			Before(func(r *http.Request) {
				r.Header.Set("Accept", "application/json, text/plain, */*")
				r.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
				r.Header.Set("User-Agent", zhttp.RandomUserAgent())
				r.Header.Set("Referer", r.URL.String())
				r.Header.Set("Origin", "https://time.geekbang.org")
			}).
			After(func(r *http.Response) error {
				if zhttp.IsHTTPSuccessStatus(r.StatusCode) {
					if stat, key, err := global.Storage.Put(destName, r.Body); err != nil {
						return err
					} else if stat.Size() <= 0 {
						return fmt.Errorf("[%s] is empty", key)
					}
					return nil
				}
				if zhttp.IsHTTPStatusSleep(r.StatusCode) {
					time.Sleep(time.Second * 10)
				}
				if zhttp.IsHTTPStatusRetryable(r.StatusCode) {
					return fmt.Errorf("http status: %s", r.Status)
				}
				return zhttp.BreakRetryError(fmt.Errorf("http status: %s", r.Status))
			}).
			DoWithRetry(retryCtx, http.MethodGet, dowloadURL, nil)
		if err != nil {
			return "", err
		}
	}

	m3u8Path := path.Join(dir, "index.m3u8")
	stat, destKey, err := global.Storage.Put(m3u8Path, io.NopCloser(&buff))
	if err != nil {
		return "", err
	} else if stat.Size() <= 0 {
		return "", fmt.Errorf("[%s] size is zero", destKey)
	}
	destDir := path.Dir(destKey)
	concatPath := path.Join(path.Dir(destDir), fmt.Sprintf("%s.mp4", fileName))
	ffmpeg_command := []string{
		"-allowed_extensions",
		"ALL",
		"-protocol_whitelist",
		"concat,file,http,https,tcp,tls,crypto",
		"-i",
		path.Join(destDir, "index.m3u8"),
		"-hls_key_info_file",
		path.Join(destDir, "key.key"),
		"-c",
		"copy",
		concatPath,
	}
	output, err := exec.CommandContext(retryCtx, "ffmpeg", ffmpeg_command...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s,%s", err.Error(), string(output))
	}
	if s, _ := os.Stat(concatPath); s != nil && s.Size() > 0 {
		_ = os.RemoveAll(destDir)
	}
	return concatPath, nil
}

func Audio(ctx context.Context, dowloadURL, dir, fileName string) (string, error) {
	retryCtx, retryCancel := context.WithTimeout(ctx, time.Minute*30)
	defer retryCancel()
	destName := path.Join(dir, fmt.Sprintf("%s.mp3", fileName))
	err := zhttp.R.
		Before(func(r *http.Request) {
			r.Header.Set("Accept", "application/json, text/plain, */*")
			r.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
			r.Header.Set("User-Agent", zhttp.RandomUserAgent())
			r.Header.Set("Referer", r.URL.String())
			r.Header.Set("Origin", "https://time.geekbang.org")
		}).
		After(func(r *http.Response) error {
			if zhttp.IsHTTPSuccessStatus(r.StatusCode) {
				if stat, key, err := global.Storage.Put(destName, r.Body); err != nil {
					return err
				} else if stat.Size() <= 0 {
					return fmt.Errorf("audio empty: [%s]", key)
				}
				return nil
			}
			if zhttp.IsHTTPStatusSleep(r.StatusCode) {
				time.Sleep(time.Second * 10)
			}
			if zhttp.IsHTTPStatusRetryable(r.StatusCode) {
				return fmt.Errorf("http status: %s", r.Status)
			}
			return zhttp.BreakRetryError(fmt.Errorf("http status: %s", r.Status))
		}).
		DoWithRetry(retryCtx, http.MethodGet, dowloadURL, nil)
	if err != nil {
		return "", err
	}
	return destName, nil
}
