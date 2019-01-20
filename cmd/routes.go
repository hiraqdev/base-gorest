package cmd

import (
	"os"

	"github.com/hiraqdev/base-gorest/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// routeCmd used to list all available routers
var routeCmd = &cobra.Command{
	Use:   "routes",
	Short: "List of available routes",
	Run: func(cmd *cobra.Command, args []string) {
		var data [][]string

		for route, controller := range app.Routers {
			r := []string{route, controller.Method}
			data = append(data, r)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Path", "HTTP Method"})

		for _, v := range data {
			table.Append(v)
		}

		table.Render()
	},
}
