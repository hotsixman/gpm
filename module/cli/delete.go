package cli

import (
	"fmt"
	"gpm/module/client"
	"gpm/module/logger"
	"gpm/module/types"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete process",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			logger.Errorln("1 argument is needed.")
			return
		}

		deleteMessage := types.DeleteMessage{
			Type: "delete",
			Name: args[0],
		}

		conn, reader, err := client.MakeUDSConn()
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		resultMessage, err := client.Delete(conn, reader, deleteMessage)
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		if resultMessage.Success {
			logger.Logln(fmt.Sprintf("Successfully deleted process \"%s\".", deleteMessage.Name))
			os.Exit(0)
		} else {
			logger.Errorln(resultMessage.Error)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
