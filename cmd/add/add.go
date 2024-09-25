package add

import (
	"database/sql"
	"fmt"
	"github.com/LuckyBOYZ/todos/repository"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "Add Command",
	Long:  `Adding a new task to your list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please provide a description for your task.")
			return
		}
		todosDatabase := repository.NewTodosDatabase()
		todoToSave := createTodo(args[0])
		err := todosDatabase.Save(todoToSave)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("task was added successfully with an id", todoToSave.Id)
	},
}

func init() {
}

func createTodo(desc string) *repository.Todo {
	return &repository.Todo{
		Description: desc,
		Done:        false,
		Created:     sql.NullTime{Time: time.Now(), Valid: true},
	}
}
