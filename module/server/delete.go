package server

import (
	"gpm/module/types"
	"gpm/module/util"
	"net"

	"github.com/mitchellh/mapstructure"
)

func (server *Server) delete(conn net.Conn, message map[string]any) error {
	var deleteMessage types.DeleteMessage
	resultMessage := types.DeleteResultMessage{
		Type:    "deleteResult",
		Success: false,
		Error:   "",
	}

	err := mapstructure.Decode(message, &deleteMessage)
	if err != nil {
		server.mainLogger.Errorln(err)
		resultMessage.Error = err.Error()
		err = util.SendMessage(conn, resultMessage)
		if err != nil {
			return err
		}
	}

	err = server.pm.Delete(deleteMessage)
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
