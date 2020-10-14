package main

import (
	"fmt"
	"os"

	"github.com/jrbeverly/bmx/config"
	"github.com/jrbeverly/bmx/console"

	"github.com/spf13/cobra"
)

var (
	userConfig config.UserConfig
	consolerw  *console.DefaultConsoleReader
)

func init() {
	userConfig = (config.ConfigLoader{}).LoadConfigs()
	consolerw = console.NewConsoleReader()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{}
