package task

type TaskListRequest struct {
	TaskPid string `json:"task_pid"  form:"task_pid"`
	Xstatus int    `json:"xstatus" form:"xstatus"`
	Page    int    `json:"page" form:"page"`
	PerPage int    `json:"perPage"  form:"perPage"`
}

type TaskListResponse struct {
	Count int64  `json:"count"`
	Rows  []Task `json:"rows"`
}

type Task struct {
	// task id
	TaskId string `json:"task_id,omitempty"`
	// task pid
	TaskPid string `json:"task_pid,omitempty"`
	// task name
	TaskName string `json:"task_name,omitempty"`
	// task type
	TaskType string `json:"task_type,omitempty"`
	// status
	Status int32 `json:"status,omitempty"`
	// Message
	Message []byte `json:"message,omitempty"`
	// statistics
	Statistics TaskStatistics `json:"statistics,omitempty"`
	// created_at
	CreatedAt int64 `json:"created_at,omitempty"`
	// updated_at
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// deleted_at
	DeletedAt int64 `json:"deleted_at,omitempty"`
}

type TaskStatistics struct {
	Count int         `json:"count"`
	Items map[int]int `json:"items"`
}

type RetryRequest struct {
	// task pid
	Pid string `json:"pid,omitempty" form:"pid"`
	// task ids
	Ids string `json:"ids,omitempty" form:"ids"`
}
