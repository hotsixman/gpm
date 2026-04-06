package server

import (
	"fmt"
	"gpm/module/util"
	"net"
	"os"
	"sync"
)

func NewUDSServer() (*Server, error) {
	socketPath := util.GetUDSPath()

	if _, err := os.Stat(socketPath); err == nil {
		os.Remove(socketPath)
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on uds: %v", err)
	}

	server := &Server{
		mainLogger: nil,
		listener:   listener,
		client:     make(map[string]*ServerSideClient),
		mutex:      &sync.Mutex{},
		pm:         nil,
	}

	go server.accept()

	return server, nil
}
