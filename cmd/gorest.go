package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addr string
var readTimeout int
var writeTimeout int

var defaultPort = ":8080"
var defaultReadTimeout = 15
var defaultWriteTimeout = 15

// gorestCmd used as main command line action
var gorestCmd = &cobra.Command{
	Use:   "gorest",
	Short: "Gorest is a skeleton rest api using golang and gorilla mux",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Thank you for using Gorest, if you want to run your engine, use server command")
	},
}

func init() {
	gorestCmd.AddCommand(serverCmd)
	gorestCmd.AddCommand(routeCmd)
}

// Execute main cmd app
func Execute() {
	if err := gorestCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
