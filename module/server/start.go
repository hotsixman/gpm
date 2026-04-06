package server

import (
	"gpm/module/types"
	"gpm/module/util"
	"net"

	"github.com/mitchellh/mapstructure"
)

func (server *Server) start(conn net.Conn, message map[string]any) error {
	var startMessage types.StartMessage
	resultMessage := types.StartResultMessage{
		Type:    "startResult",
		Success: false,
		Error:   "",
	}

	err := mapstructure.Decode(message, &startMessage)
	if err != nil {
		server.mainLogger.Errorln(err)
		resultMessage.Error = err.Error()
		err = util.SendMessage(conn, resultMessage)
		if err != nil {
			return err
		}
	}

	err = server.pm.Start(startMessage)
	if err != nil {
		server.mainLogger.Errorln(err)
		resultMessage.Error = err.Error()
		err = util.SendMessage(conn, resultMessage)
		if err != nil {
			return err
		}
	}

	resultMessage.Success = true
	err = util.SendMessage(conn, resultMessage)
	if err != nil {
		return err
	}
	return nil
}
