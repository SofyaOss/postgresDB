package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewDb(ConnString string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), ConnString)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) AllTasks() ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT * FROM tasks ORDER BY id;`,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.Id,
			&t.Opened,
			&t.Closed,
			&t.AuthorId,
			&t.AssignedId,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

func (s *Storage) TasksAuthor(id int) ([]Task, error) {
	var tasksList []Task
	rows, err := s.db.Query(context.Background(), `SELECT * FROM tasks WHERE author_id = $1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Opened, &t.Closed, &t.AuthorId,
			&t.AssignedId, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		tasksList = append(tasksList, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasksList, nil
}

func (s *Storage) TasksLabel(label string) ([]Task, error) {
	var tasksList []Task
	rows, err := s.db.Query(context.Background(), `SELECT * FROM tasks_labels
	join tasks ON tasks.id = tasks_labels.task_id
	join labels ON tasks_labels.label_id = labels.id and labels.name = $1;`, label)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Opened, &t.Closed, &t.AuthorId,
			&t.AssignedId, &t.Title, &t.Content)
		if err != nil {
			return nil, err
		}
		tasksList = append(tasksList, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasksList, nil
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}

func (s *Storage) UpdateTask(id int, title, content string) error {
	_, err := s.db.Exec(context.Background(), `
		UPDATE tasks SET (title, content)
		VALUES ($1, $2) WHERE id = $3;
		`,
		title,
		content,
		id,
	)
	return err
}

func (s *Storage) RmTask(id int) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks
		WHERE id = $1;
		`,
		id,
	)
	return err
}
