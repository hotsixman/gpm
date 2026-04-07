package client

import (
	"bufio"
	"gpm/module/types"
	"gpm/module/util"
	"net"
	"strings"
)

func Delete(conn net.Conn, reader *bufio.Reader, deleteMessage types.DeleteMessage) (message *types.DeleteResultMessage, err error) {
	err = util.SendMessage(conn, deleteMessage)
	if err != nil {
		return nil, err
	}

	messageJSON, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	message, err = util.ParseMessage[types.DeleteResultMessage]([]byte(strings.TrimSpace(messageJSON)))
	if err != nil {
		return nil, err
	}

	return message, nil
}
