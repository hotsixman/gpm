package cli

import (
	"bufio"
	"gpm/module/client"
	"gpm/module/logger"
	"os"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect [name]",
	Short: "Connect to process",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			logger.Errorln("1 argument is needed.")
			return
		}

		closeChan := make(chan bool)
		conn, reader, err := client.MakeUDSConn()
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		client, err := client.NewClient(args[0], conn, reader, closeChan)
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		// stdin 에서 command를 읽어서 전송
		scanner := bufio.NewScanner(os.Stdin)
		go func() {
			for scanner.Scan() {
				command := scanner.Text()
				client.Command(command)
			}
		}()

		<-closeChan
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
