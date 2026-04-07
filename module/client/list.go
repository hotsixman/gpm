package client

import (
	"bufio"
	"gpm/module/types"
	"gpm/module/util"
	"net"
	"strings"
)

func List(conn net.Conn, reader *bufio.Reader) (message *types.ListResultMessage, err error) {
	err = util.SendMessage(conn, types.ListMessage{Type: "list"})
	if err != nil {
		return nil, err
	}

	messageJSON, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	message, err = util.ParseMessage[types.ListResultMessage]([]byte(strings.TrimSpace(messageJSON)))
	if err != nil {
		return nil, err
	}

	return message, nil
}
