package task

type TaskManager struct {
	// Tasker 任务管理接口
	Tasker Tasker
	// TaskExecutor 任务执行接口
	TaskExecutor TaskExecutor
}

// NewTaskManager 创建任务管理器
func NewTaskManager(tasker Tasker, taskExecutor TaskExecutor) *TaskManager {
	return &TaskManager{
		Tasker:       tasker,
		TaskExecutor: taskExecutor,
	}
}

// CreateTask 创建任务
func (m *TaskManager) CreateTask(task *TaskInfo) error {
	return m.Tasker.CreateTask(task)
}

// GetTask 获取任务
func (m *TaskManager) GetTask(id string) (*TaskInfo, error) {
	return m.Tasker.GetTask(id)
}

// UpdateTask 更新任务
func (m *TaskManager) UpdateTask(task *TaskInfo) error {
	return m.Tasker.UpdateTask(task)
}

// DeleteTask 删除任务
func (m *TaskManager) DeleteTask(id string) error {
	return m.Tasker.DeleteTask(id)
}

// ExecuteTask 执行任务
func (m *TaskManager) ExecuteTask(task *TaskInfo) error {
	return m.TaskExecutor.ExecuteTask(task)
}
