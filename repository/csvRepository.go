package repository

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

type TodosCsv struct{}

func newTodosCsv() *TodosCsv {
	return &TodosCsv{}
}

func (t *TodosCsv) Save(data *Todo) error {
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
		return err
	}
	addIdToTodo(data, records)
	return writer.Write(todoToStringArray(data))
}

func (t *TodosCsv) FindById(id int) (*Todo, error) {
	f, err := openCSVFile()
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, v := range records {
		idToCheck, err := strconv.Atoi(v[0])
		if err != nil {
			return nil, err
		}
		if idToCheck == id {
			return stringToTodoArr(v), nil
		}
	}
	return nil, fmt.Errorf("todo with an id %d doesn,t exist", id)
}

func (t *TodosCsv) FindAll() ([]Todo, error) {
	f, err := openCSVFile()
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return convertStringArrToTodoArr(records), nil
}

func (t *TodosCsv) FindAllNotFinishedTodos() ([]Todo, error) {
	f, err := openCSVFile()
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return convertStringArrToTodoArr(records), nil
}

func (t *TodosCsv) Delete(id int) error {
	f, err := openCSVFile()
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(f)
	writer := csv.NewWriter(f)
	defer writer.Flush()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, v := range records {
		strId, err := strconv.Atoi(v[0])
		if err != nil {
			return err
		}
		if strId == id {
			continue
		}
		if err := writer.Write(v); err != nil {
			return err
		}
	}
	return nil
}

func (t *TodosCsv) CompleteTodo(id int, done bool) (int64, error) {
	panic("implement me")
}

func openCSVFile() (*os.File, error) {
	csvFilePath := viper.GetString("csvRepoFilePath")
	path, err := createFilePath(csvFilePath)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open a file for reading, writting or appending")
	}

	if err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, fmt.Errorf("cannot lock the file")
	}
	return f, nil
}

func createFilePath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, path[1:]), nil
	}
	return path, nil
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

var _ ITodoRepository[Todo] = (*TodosCsv)(nil)
