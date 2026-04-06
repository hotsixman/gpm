package client

import (
	"bufio"
	"gpm/module/util"
	"net"
)

func MakeUDSConn() (conn net.Conn, bufReader *bufio.Reader, err error) {
	socketPath := util.GetUDSPath()
	conn, err = net.Dial("unix", socketPath)
	if err != nil {
		return nil, nil, err
	}
	return conn, bufio.NewReader(conn), nil
}
