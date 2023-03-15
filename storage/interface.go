package storage

type InterfaceDB interface {
	NewTask(t Task) (int, error)
	AllTasks() ([]Task, error)
	UpdateTask(id int, title, content string) error
	RmTask(id int) error
	TasksAuthor(id int) ([]Task, error)
	TasksLabel(label string) ([]Task, error)
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
