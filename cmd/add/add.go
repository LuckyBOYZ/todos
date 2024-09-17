package add

import (
	"fmt"
	"github.com/LuckyBOYZ/todos/cmd/db"
	"github.com/spf13/cobra"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "Add Command",
	Long:  `Adding a new task to your list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a description for your task.")
			return
		}
		db.AddNewTodo(createTodo(args[0]))
	},
}

func init() {
}

func createTodo(desc string) *db.Todo {
	return &db.Todo{
		Description: desc,
		Done:        false,
		Created:     time.Now(),
	}
}
