package cmd

import (
	"fmt"
	"os"

	"github.com/hiraqdev/base-gorest/app"
	"github.com/spf13/cobra"
)

var addr string
var readTimeout int
var writeTimeout int

var defaultPort = ":8080"
var defaultReadTimeout = 15
var defaultWriteTimeout = 15

var serverCmd = &cobra.Command{
	Use:   "gorest",
	Short: "Gorest is a skeleton rest api using golang and gorilla",
	Run: func(cmd *cobra.Command, args []string) {
		options := app.ServerConfig{
			Address:      addr,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		}

		app.Env()
		app.Gorest(options)
	},
}

func init() {
	serverCmd.Flags().StringVarP(&addr, "addr", "", defaultPort, "Setup your ip address and port, default: 8080")
	serverCmd.Flags().IntVar(&writeTimeout, "writeTimeout", defaultWriteTimeout, "Set your write timeout, default: 15 seconds")
	serverCmd.Flags().IntVar(&readTimeout, "readTimeout", defaultReadTimeout, "Set your read timeout, default: 15 seconds")
}

// Execute main cmd app
func Execute() {
	if err := serverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
