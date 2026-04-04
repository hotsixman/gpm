package daemon

import (
	"gpm/module/database"
	"gpm/module/logger"
	"gpm/module/uds"
	"os"
	"os/exec"
	"time"
)

const DAEMON_ENV = "GPM_DAEMON_PROCESS"

func Daemonize() {
	if os.Getenv(DAEMON_ENV) == "1" {
		return
	}

	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Env = append(os.Environ(), DAEMON_ENV+"=1")

	setupDaemon(cmd)
	if err := cmd.Start(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func DaemonInit() {
	if os.Getenv(DAEMON_ENV) != "1" {
		return
	}

	// main logger
	log, err := logger.GetMainLogger()
	if err != nil {
		os.Exit(1)
	}

	// db
	db, err := database.OpenDB()
	if err != nil {
		log.Log("Cannot open db.")
		os.Exit(1)
	}
	defer db.Close()

	//
	_, running, err := PIDManager.CheckPID()
	if err != nil {
		log.Log("Cannot check GPM daemon is running.")
		os.Exit(1)
	}
	if running {
		log.Log("GPM is already running.")
		os.Exit(1)
	}

	//
	err = PIDManager.RecordPid()
	if err != nil {
		log.Log("Cannot record pid.")
		os.Exit(1)
	}
	defer PIDManager.DeletePid()

	//
	udsServer, err := uds.Listen()
	if err != nil {
		log.Log("Cannot listen uds server.")
		os.Exit(1)
	}
	log.SetUDSServer(udsServer)

	//
	go func() {
		for {
			log.Log(time.Now())
			time.Sleep(time.Second)
		}
	}()

	select {}
}
