package client

import (
	"bufio"
	"errors"
	"gpm/module/logger"
	"gpm/module/types"
	"gpm/module/util"
	"io"
	"net"
)

type Client struct {
	conn   net.Conn
	reader *bufio.Reader
}

func NewClient(name string, conn net.Conn, reader *bufio.Reader, closeChan chan bool) (*Client, error) {
	client := &Client{
		conn:   conn,
		reader: reader,
	}

	// Send connection request message
	message := types.ConnectRequestMessage{
		Type: "connect",
		Name: name,
	}
	err := util.SendMessage(conn, message)
	if err != nil {
		return nil, err
	}

	// @todo Need to check ConnectResponseMessage

	go func() {
		// close
		defer func() {
			closeChan <- true
			close(closeChan)
		}()
		// log
		for {
			messageJSON, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					logger.Logln("Connection Closed.")
				} else {
					logger.Errorln(err)
				}
				return
			}

			message, err := util.ParseMessage[map[string]string]([]byte(messageJSON))
			if err != nil {
				logger.Errorln(err)
				continue
			}

			switch (*message)["type"] {
			case "log":
				if (*message)["message"] != "" {
					logger.Logln((*message)["message"])
				}
			case "error":
				if (*message)["message"] != "" {
					logger.Errorln((*message)["message"])
				}
			}
		}
	}()

	return client, nil
}

func (client *Client) Command(command string) error {
	message := types.CommandMessage{
		Type:    "command",
		Command: command,
	}

	return util.SendMessage(client.conn, message)
}
