package util

import (
	"encoding/json"
	"net"
)

func SendMessage(conn net.Conn, message any) error {
	JSON, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(append(JSON, '\n')))
	return err
}

func ParseMessage[T any](JSON []byte) (*T, error) {
	var message T
	err := json.Unmarshal(JSON, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
