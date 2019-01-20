package cmd

import (
	"github.com/hiraqdev/base-gorest/app"
	"github.com/spf13/cobra"
)

// serverCmd used to run api engine
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run your rest api engine",
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
