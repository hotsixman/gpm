package uds

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gpm/module/logger"
	"log"
	"net"
	"strings"
)

type UDSClient struct {
	conn net.Conn
}

func Connect() (*UDSClient, error) {
	socketPath := GetSocketPath()
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to daemon: %v", err)
	}

	client := &UDSClient{
		conn: conn,
	}

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Println(err)
				return
			}

			var data any
			err = json.Unmarshal([]byte(strings.TrimSpace(message)), &data)
			if err == nil {
				logger.Logln(data)
			} else {
				logger.Errorln(message)
			}
		}
	}()

	return client, nil
}
