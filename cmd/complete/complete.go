package complete

import (
	"fmt"
	"github.com/LuckyBOYZ/todos/repository"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete your task by given id",
	Long:  "Complete your task by given id",
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
			repository := repository.GetTodoRepository()
			undone, _ := cmd.Flags().GetBool("undone")
			lastInsertedId, err := repository.CompleteTodo(todoId, !undone)
			if err != nil {
				fmt.Println("Something went wrong during updating task status", err)
				os.Exit(1)
			}
			if lastInsertedId < 1 {
				fmt.Printf("No task found for given id %d\n", todoId)
				return
			}

			if !undone {
				fmt.Printf("Task with id %d is completed successfully!\n", todoId)
			} else {
				fmt.Printf("Task with id %d is back to incomplete status!\n", todoId)
			}
		}
	},
}

func init() {
	Cmd.Flags().BoolP("undone", "u", false, "change task status to incomplete")
}
