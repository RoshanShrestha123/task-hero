package helper

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Task struct {
	Id          int
	Task        string
	IsCompleted int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Add(db *sql.DB, title string) {
	var task Task = Task{Task: strings.TrimSpace(title), IsCompleted: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	insertQuery := `insert into todos (task, isCompleted, createdAt, updatedAt) values (?, ? , ?, ?)`
	_, err := db.Exec(insertQuery, task.Task, task.IsCompleted, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		panic(err)
	}

}

func List(db *sql.DB) []Task {
	var tasks []Task
	selectQuery := `select * from todos`
	rows, err := db.Query(selectQuery)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var _task Task
		err = rows.Scan(&_task.Id, &_task.Task, &_task.IsCompleted, &_task.CreatedAt, &_task.UpdatedAt)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, _task)
	}

	return tasks
}

func Delete(db *sql.DB, id int) {
	_, err := db.Exec("Delete from todos where id=?", id)
	if err != nil {
		panic(err)
	}

	fmt.Sprintf("%v has been deleted", id)
}
