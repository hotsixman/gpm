package client

import (
	"bufio"
	"gpm/module/types"
	"gpm/module/util"
	"net"
	"strings"
)

func Stop(conn net.Conn, reader *bufio.Reader, stopMessage types.StopMessage) (message *types.StopResultMessage, err error) {
	err = util.SendMessage(conn, stopMessage)
	if err != nil {
		return nil, err
	}

	messageJSON, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	message, err = util.ParseMessage[types.StopResultMessage]([]byte(strings.TrimSpace(messageJSON)))
	if err != nil {
		return nil, err
	}

	return message, nil
}
