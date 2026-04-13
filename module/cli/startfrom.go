package cli

import (
	"encoding/json"
	"fmt"
	"geep/module/client"
	"geep/module/logger"
	"geep/module/types"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type startData struct {
	Name            *string           `json:"name"`
	Run             *string           `json:"run"`
	Args            []string          `json:"args"`
	Env             map[string]string `json:"env"`
	Cwd             *string           `json:"cwd"`
	MaxRecoverCount *int              `json:"maxRecoverCount"`
	MaxLogfileSize  *int              `json:"maxLogfileSize"`
}

var startfromCommand = &cobra.Command{
	Use:   "startfrom [cmd]",
	Short: "Start a new process from json",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jsonPath := args[0]
		data, err := loadJson(jsonPath)
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		startMessage := types.StartMessage{Type: "start"}
		if data.Name == nil {
			logger.Errorln(types.NoNameError{JsonPath: jsonPath})
			os.Exit(1)
		} else {
			startMessage.Name = *data.Name
		}
		if data.Run == nil {
			logger.Errorln(types.NoRunError{JsonPath: jsonPath})
			os.Exit(1)
		} else {
			startMessage.Run = *data.Run
		}
		if data.Args == nil {
			startMessage.Args = make([]string, 0)
		} else {
			startMessage.Args = data.Args
		}
		if data.Env == nil {
			startMessage.Env = make(map[string]string)
		} else {
			startMessage.Env = data.Env
		}
		if data.MaxRecoverCount == nil {
			startMessage.MaxRecoverCount = 10
		} else {
			startMessage.MaxRecoverCount = *data.MaxRecoverCount
		}
		if data.MaxLogfileSize == nil {
			startMessage.MaxLogfileSize = 1024 * 100
		} else {
			startMessage.MaxLogfileSize = *data.MaxLogfileSize
		}

		conn, reader, err := client.MakeUDSConn()
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		resultMessage, err := client.Start(conn, reader, startMessage)
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		if resultMessage.Success {
			logger.Logln(fmt.Sprintf("Successfully started process \"%s\".", startMessage.Name))
			os.Exit(0)
		} else {
			logger.Errorln(resultMessage.Error)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(startfromCommand)
}

func loadJson(jsonPath string) (*startData, error) {
	jsonPath = filepath.Clean(jsonPath)
	if !filepath.IsAbs(jsonPath) {
		cwd, err := os.Getwd()
		if err == nil {
			jsonPath = filepath.Join(cwd, jsonPath)
		}
	}

	_, err := os.Stat(jsonPath)
	if err != nil {
		return nil, err
	}

	jsonByte, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	data := &startData{}
	err = json.Unmarshal(jsonByte, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
