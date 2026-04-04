package cli

import (
	"gpm/module/daemon"
	"log"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminate GPM daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		pid, running, err := daemon.PIDManager.CheckPID()
		if err != nil {
			log.Println("???")
			os.Exit(1)
		}

		if !running {
			log.Println("GPM daemon is not running.")
			os.Exit(1)
		}

		err = syscall.Kill(pid, syscall.SIGTERM)
		if err != nil {
			log.Println("Cannot kill GPM daemon.")
			os.Exit(1)
		}
		log.Println("Killed GPM daemon.")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(terminateCmd)
}
