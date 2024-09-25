package repository

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

type Todo struct {
	Id          int
	Description string
	Done        bool
	Created     sql.NullTime
}

func AddNewTodo(t *Todo) {
	f, err := openCSVFile()
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(f)
	defer func(file *os.File) {
		_ = file.Close()
	}(f)
	defer func(file *os.File) {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	}(f)
	defer writer.Flush()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	addIdToTodo(t, records)
	if err := writer.Write(todoToStringArray(*t)); err != nil {
		panic(err)
	}
}

func GetTodos(all bool) [][]string {
	f, err := openCSVFile()
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	if all {
		return records
	}
	var undoneTodos [][]string
	for _, r := range records {
		isUndone := r[2]
		if isUndone == "false" {
			undoneTodos = append(undoneTodos, r)
		}
	}
	return undoneTodos
}

func DeleteTodoById(id int) {
	f, err := openCSVFile()
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(f)
	writer := csv.NewWriter(f)
	defer writer.Flush()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	for _, v := range records {
		strId, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}
		if strId == id {
			continue
		}
		if err := writer.Write(v); err != nil {
			panic(err)
		}
	}
}

func openCSVFile() (*os.File, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot get home directory")
	}
	filePath := filepath.Join(homeDir, "todos.csv")
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open a file for reading")
	}

	if err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("cannot lock the file")
	}
	return f, nil
}

func todoToStringArray(t Todo) []string {
	var epoch int64
	if t.Created.Valid {
		epoch = t.Created.Time.Unix()
	}
	return []string{
		strconv.Itoa(t.Id),
		t.Description,
		strconv.FormatBool(t.Done),
		strconv.FormatInt(epoch, 10),
	}
}

func addIdToTodo(t *Todo, records [][]string) {
	if t.Id < 1 {
		t.Id = getNextIdFromCSV(records)
	}
}

func getNextIdFromCSV(records [][]string) int {
	var lastId int
	for _, r := range records {
		id, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}
		if id > lastId {
			lastId = id
		}
	}
	return lastId + 1
}
