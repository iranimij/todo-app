package todo

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize opens the database and creates the table if not exists
func Initialize() error {
	var err error
	db, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT 0
	);`
	_, err = db.Exec(createTable)
	return err
}

// AddTask inserts a new todo into the database
func AddTask(task string) error {
	_, err := db.Exec("INSERT INTO todos (task, done) VALUES (?, 0)", task)
	return err
}

// GetAll retrieves all todos from the database
func GetAll() ([]Todo, error) {
	rows, err := db.Query("SELECT id, task, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err = rows.Scan(&t.ID, &t.Task, &t.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

// MarkDone updates a todo’s done status
func MarkDone(id int) error {
	_, err := db.Exec("UPDATE todos SET done = 1 WHERE id = ?", id)
	return err
}

// Delete removes a todo by ID
func Delete(id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
