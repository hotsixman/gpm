package cli

import (
	"fmt"
	"geep/module/client"
	"geep/module/logger"
	"geep/module/util"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List processes",
	Run: func(cmd *cobra.Command, args []string) {
		conn, reader, err := client.MakeUDSConn()
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		resultMessage, err := client.List(conn, reader)
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		table := util.ListTable(*resultMessage)
		fmt.Println(table)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
