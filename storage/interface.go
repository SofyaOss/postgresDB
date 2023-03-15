package storage

type InterfaceDB interface {
	NewTask(t Task) (int, error)
	Tasks(taskID, authorID int) ([]Task, error)
	UpdateTask(id int, task Task) (int, error)
	RmTaskId(id int) (int, error)
}

type User struct {
	id   int
	name string
}

type Label struct {
	id   int
	name string
}

type TaskLabel struct {
	taskId  int
	labelId int
}

type Task struct {
	Id         int
	Opened     int
	Closed     int
	AuthorId   int
	AssignedId int
	Title      string
	Content    string
}
