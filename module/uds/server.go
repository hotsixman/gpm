package uds

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/google/uuid"
)

type UDSServer struct {
	listener net.Listener
	clients  map[string]net.Conn
	mutex    *sync.Mutex
}

func Listen() (*UDSServer, error) {
	socketPath := GetSocketPath()

	if _, err := os.Stat(socketPath); err == nil {
		os.Remove(socketPath)
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on uds: %v", err)
	}

	server := &UDSServer{
		listener: listener,
		clients:  make(map[string]net.Conn),
		mutex:    &sync.Mutex{},
	}

	server.accept()

	return server, nil
}

func (this *UDSServer) Broadcast(JSON []byte) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for _, conn := range this.clients {
		go func() {
			conn.Write(append(JSON, '\n'))
		}()
	}
}

func (this *UDSServer) accept() {
	go func() {
		for {
			conn, err := this.listener.Accept()
			if err != nil {
				continue
			}

			go this.handleClient(conn)
		}
	}()
}

func (this *UDSServer) handleClient(conn net.Conn) {
	this.mutex.Lock()
	id := ""
	for {
		id = uuid.New().String()
		if this.clients[id] == nil {
			break
		}
	}
	this.clients[id] = conn
	this.mutex.Unlock()

	reader := bufio.NewReader(conn)

	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			this.mutex.Lock()
			delete(this.clients, id)
			this.mutex.Unlock()
			return
		}
	}
}
