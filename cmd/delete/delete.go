package delete

import (
	"fmt"
	"github.com/LuckyBOYZ/todos/cmd/db"
	"github.com/spf13/cobra"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your task by given id",
	Long:  "Delete your task by given id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide id")
			return
		}
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Given id is not parsable")
			return
		}
		db.DeleteTodoById(int(id))
	},
}

func init() {
}
