package server

import (
	"gpm/module/types"
	"gpm/module/util"
	"net"

	"github.com/mitchellh/mapstructure"
)

func (server *Server) stop(conn net.Conn, message map[string]any) error {
	var stopMessage types.StopMessage
	resultMessage := types.StopResultMessage{
		Type:    "stopResult",
		Success: false,
		Error:   "",
	}

	err := mapstructure.Decode(message, &stopMessage)
	if err != nil {
		server.mainLogger.Errorln(err)
		resultMessage.Error = err.Error()
		err = util.SendMessage(conn, resultMessage)
		if err != nil {
			return err
		}
	}

	err = server.pm.Stop(stopMessage)
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
