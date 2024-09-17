package cmd

import (
	"github.com/LuckyBOYZ/todos/cmd/add"
	"github.com/LuckyBOYZ/todos/cmd/complete"
	"github.com/LuckyBOYZ/todos/cmd/delete"
	"github.com/LuckyBOYZ/todos/cmd/list"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Application to manage your tasks",
	Long:  `You can add, delete, complete and list your tasks using this application`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(add.Cmd)
	rootCmd.AddCommand(list.Cmd)
	rootCmd.AddCommand(delete.Cmd)
	rootCmd.AddCommand(complete.Cmd)
}
