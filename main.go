package main

import (
	"fmt"
	"github.com/LuckyBOYZ/todos/cmd"
	"github.com/LuckyBOYZ/todos/configuration"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func main() {
	pflag.BoolP("csv", "c", false,
		`whether to use csv repository. Default: false.
If not passed, then postgres database is used`)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := configuration.LoadConfiguration(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.Execute()
}
