package main

import (
	"fmt"
	"github.com/LuckyBOYZ/todos/cmd"
	"github.com/LuckyBOYZ/todos/configuration"
)

func main() {
	if err := configuration.LoadConfiguration(); err != nil {
		fmt.Println(err)
		return
	}
	cmd.Execute()
}
