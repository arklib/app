package base

type Tasks struct {
	SyncERPUser any
}

type Task struct {
	Name    string
	Handler func()

	Driver  any
	Timeout any
	Push    func()
	Pull    func()
	Ack     func()
}

func (base *Base) GetTasks() *Tasks {
	if base.Tasks != nil {
		return base.Tasks
	}

	base.Tasks = &Tasks{
		SyncERPUser: &Task{
			Driver:  "",
			Name:    "sync_erp_user",
			Timeout: 60,
		},
	}
	return base.Tasks
}
