package delete

import (
	"fmt"
	"github.com/LuckyBOYZ/todos/repository"
	"github.com/spf13/cobra"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your todo by given id",
	Long:  "Delete your todo by given id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide task id")
		} else if len(args) > 1 {
			fmt.Println("Please provide only one task id")
		} else {
			todoId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Please provide valid task id")
				return
			}
			todosDatabase := repository.NewTodosDatabase()
			err = todosDatabase.Delete(todoId)
			if err != nil {
				fmt.Println("something went wrong during updating task status.", err)
			} else {
				fmt.Printf("Task with id %d was deleted succesfully!\n", todoId)
			}
		}
	},
}

func init() {
}
