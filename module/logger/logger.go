package logger

import (
	"encoding/json"
	"fmt"
	"gpm/module/util"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Logln(v ...any) {
	message := strings.TrimRight(fmt.Sprintln(v...), " \t\n\r")
	timeString := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	header := "\033[32m" + timeString + " [LOG]" + "\033[0m"
	fmt.Println(header, message)
}

func Errorln(v ...any) {
	message := strings.TrimRight(fmt.Sprintln(v...), " \t\n\r")
	timeString := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	header := "\033[31m" + timeString + " [ERROR]" + "\033[0m"
	fmt.Println(header, message)
}

type Broadcastable interface {
	Broadcast(JSON []byte)
}

type Logger struct {
	dirPath        string
	name           string
	timeRecorded   bool
	errorSeperated bool
	udsServer      Broadcastable
}

func GetMainLogger() (*Logger, error) {
	homeDir, err := util.GetHomeDirPath()
	if err != nil {
		return nil, err
	}

	dirPath := filepath.Join(homeDir, "log")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &Logger{
		dirPath,
		"",
		true,
		false,
		nil,
	}, nil
}

func GetLogger(name string, timeRecorded bool, errorSeperated bool, broadCastable Broadcastable) (*Logger, error) {
	homeDir, err := util.GetHomeDirPath()
	if err != nil {
		return nil, err
	}

	dirPath := filepath.Join(homeDir, "log-process", filepath.Clean(name))
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &Logger{
		dirPath,
		name,
		timeRecorded,
		errorSeperated,
		broadCastable,
	}, nil
}

func (this *Logger) SetUDSServer(udsServer Broadcastable) {
	this.udsServer = udsServer
}

func (this *Logger) Log(v ...any) {
	message := strings.TrimRight(fmt.Sprintln(v...), " \t\n\r")
	log.Println(message)

	if this.udsServer != nil {
		JSON, err := json.Marshal(message)
		if err == nil {
			this.udsServer.Broadcast(JSON)
		}
	}
}
