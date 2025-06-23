package serve

import (
	"github.com/rmhyde/fusion/cmd/utils"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "CLI Tool for combining json files within the current or specified folder",
	RunE:  runE,
}

func init() {
	ServeCmd.PersistentFlags().Int("port", 8080, "Port number of the webserver, will default to 8080")
	ServeCmd.PersistentFlags().String("host", "localhost", "hostname to list")
}

func runE(cmd *cobra.Command, args []string) error {
	options, err := utils.NewServeOptions(cmd, args)
	if err != nil {
		return err
	}

	return options.StartWebServer()
}
