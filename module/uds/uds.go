package uds

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func GetSocketPath() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".gpm")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	return filepath.Join(dir, "gpm.sock")
}

func Listen() (net.Listener, error) {
	socketPath := GetSocketPath()

	if _, err := os.Stat(socketPath); err == nil {
		os.Remove(socketPath)
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on uds: %v", err)
	}

	// 모든 사용자가 접근 가능하도록 권한 설정
	// os.Chmod(socketPath, 0666)

	return listener, nil
}

func Connect() (net.Conn, error) {
	socketPath := GetSocketPath()
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to daemon: %v", err)
	}
	return conn, nil
}
