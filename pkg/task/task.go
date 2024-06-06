package task

type TaskInfo struct {
	// ID 任务ID
	ID string
	// Name 任务名称
	Name string
	// Type 任务类型
	Type string

	// Status 任务状态
	Status string

	// StartTime 任务开始时间
	StartTime string

	// EndTime 任务结束时间
	EndTime string

	// 任务更新时间
	UpdateTime string

	// Payload 任务数据
	Payload []byte

	// Result 任务结果
	Result []byte
}

type Tasker interface {
	// CreateTask 创建任务
	CreateTask(task *TaskInfo) error

	// GetTask 获取任务
	GetTask(id string) (*TaskInfo, error)

	// UpdateTask 更新任务
	UpdateTask(task *TaskInfo) error

	// DeleteTask 删除任务
	DeleteTask(id string) error
}

type TaskExecutor interface {
	// ExecuteTask 执行任务
	ExecuteTask(task *TaskInfo) error
}
