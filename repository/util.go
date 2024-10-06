package repository

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"time"
)

const DATE_FORMAT = "2006-01-02 15:04:05"

func GetTodoRepository() ITodoRepository[Todo] {
	useCsv, err := strconv.ParseBool(viper.GetString("csv"))
	if err != nil || !useCsv {
		return newTodosDatabase()
	} else {
		return newTodosCsv()
	}
}

func ConvertTodoArrToStringArr(todos []Todo) [][]string {
	var stringArr [][]string
	for _, v := range todos {
		stringArr = append(stringArr, todoToStringArray(&v))
	}
	return stringArr
}

func convertStringArrToTodoArr(stringArr [][]string) []Todo {
	var todoArr []Todo
	for _, v := range stringArr {
		todoArr = append(todoArr, *stringToTodoArr(v))
	}
	return todoArr
}

func stringToTodoArr(todoAsArr []string) *Todo {
	id, err := strconv.Atoi(todoAsArr[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	done, err := strconv.ParseBool(todoAsArr[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	epochInt, err := strconv.ParseInt(todoAsArr[3], 10, 64)
	createdTime := time.Unix(epochInt, 0)
	return &Todo{
		Id:          id,
		Description: todoAsArr[1],
		Done:        done,
		Created:     sql.NullTime{Time: createdTime, Valid: true},
	}
}

func CreateTodoByDescription(desc string) *Todo {
	return &Todo{
		Description: desc,
		Done:        false,
		Created:     sql.NullTime{Time: time.Now(), Valid: true},
	}
}

func todoToStringArray(t *Todo) []string {
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
