package cli

import (
	"gpm/module/uds"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to GPM daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := uds.Connect()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		select {}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
