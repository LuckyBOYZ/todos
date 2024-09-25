package repository

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
)

var (
	once         sync.Once
	loadError    error
	dbConnection *sql.DB
)

type TodosDatabase struct{}

func NewTodosDatabase() *TodosDatabase {
	once.Do(func() {
		dbConn, err := NewDatabaseConnection()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		dbConnection = dbConn
		return
	})
	if loadError != nil {
		fmt.Println(loadError.Error())
		os.Exit(1)
	}
	return &TodosDatabase{}
}

func (t *TodosDatabase) Save(data *Todo) error {
	query := `
        INSERT INTO todos (description, done, created) 
        VALUES ($1, $2, $3)
        RETURNING id
    `
	err := dbConnection.QueryRow(query, data.Description, data.Done, data.Created).Scan(&data.Id)
	if err != nil {
		return fmt.Errorf("failed to save todo. %v", err)
	}
	return nil
}

func (t *TodosDatabase) FindById(id int) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TodosDatabase) FindAll() ([]Todo, error) {
	rows, err := dbConnection.Query("SELECT id, description, done, created FROM todos")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	// Iterowanie po wierszach wyników zapytania
	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Description, &todo.Done, &todo.Created); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	// Sprawdzenie błędów podczas iteracji
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodosDatabase) Delete(id int) error {
	_, err := dbConnection.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

func (t *TodosDatabase) CompleteTodo(id int, done bool) (int64, error) {
	result, err := dbConnection.Exec("UPDATE todos SET done = $1 WHERE id = $2", done, id)
	if err != nil {
		return -1, err
	}
	updatedRows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return updatedRows, nil
}

func (t *TodosDatabase) FindAllNotFinishedTodos() ([]Todo, error) {
	rows, err := dbConnection.Query("SELECT id, description, done, created FROM todos WHERE done = false")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	// Iterowanie po wierszach wyników zapytania
	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Description, &todo.Done, &todo.Created); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	// Sprawdzenie błędów podczas iteracji
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

var _ IDatabase[Todo] = (*TodosDatabase)(nil)
