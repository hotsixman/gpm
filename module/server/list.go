package server

import (
	"gpm/module/types"
	"gpm/module/util"
	"net"

	"github.com/mitchellh/mapstructure"
)

func (server *Server) list(conn net.Conn, message map[string]any) error {
	var listMessage types.ListMessage
	resultMessage := types.ListResultMessage{
		Type: "listResult",
		List: make([]types.ListElement, 0),
	}

	err := mapstructure.Decode(message, &listMessage)
	if err != nil {
		server.mainLogger.Errorln(err)
		err = util.SendMessage(conn, resultMessage)
		if err != nil {
			return err
		}
	}

	list := server.pm.List()
	resultMessage.List = list
	err = util.SendMessage(conn, resultMessage)
	if err != nil {
		return err
	}
	return nil
}
