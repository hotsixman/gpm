package daemon

import (
	"os"
	"os/exec"
)

const envVar = "GPM_DAEMON_PROCESS"

func Daemonize() {
	if os.Getenv(envVar) == "1" {
		return
	}

	cmd := exec.Command(os.Args[0], os.Args[1:]...)

	cmd.Env = append(os.Environ(), envVar+"=1")

	setupDaemon(cmd)

	if err := cmd.Start(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
