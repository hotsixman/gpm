package cli

import (
	"fmt"
	"gpm/module/client"
	"gpm/module/logger"
	"gpm/module/types"
	"os"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [name]",
	Short: "Stop process",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			logger.Errorln("1 argument is needed.")
			return
		}

		stopMessage := types.StopMessage{
			Type: "stop",
			Name: args[0],
		}

		conn, reader, err := client.MakeUDSConn()
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		resultMessage, err := client.Stop(conn, reader, stopMessage)
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		if resultMessage.Success {
			logger.Logln(fmt.Sprintf("Successfully stopped process \"%s\".", stopMessage.Name))
			os.Exit(0)
		} else {
			logger.Errorln(resultMessage.Error)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
