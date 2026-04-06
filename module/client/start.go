package client

import (
	"bufio"
	"gpm/module/types"
	"gpm/module/util"
	"net"
	"strings"
)

func Start(conn net.Conn, reader *bufio.Reader, startMessage types.StartMessage) (message *types.StartResultMessage, err error) {
	err = util.SendMessage(conn, startMessage)
	if err != nil {
		return nil, err
	}

	messageJSON, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	message, err = util.ParseMessage[types.StartResultMessage]([]byte(strings.TrimSpace(messageJSON)))
	if err != nil {
		return nil, err
	}

	return message, nil
}
