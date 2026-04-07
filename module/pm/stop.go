package pm

import (
	"gpm/module/types"
)

func (pm *PM) Stop(message types.StopMessage) error {
	process := pm.process[message.Name]
	if process == nil {
		return &types.NoProcessError{Name: message.Name}
	}
	if process.status == "running" {
		process.status = "stop'"
		err := process.cmd.Process.Kill()
		if err != nil {
			pm.mainLogger.Logln("Cannot stop process: ", message.Name)
			return err
		}
		process.status = "stop"
	}
	pm.mainLogger.Logln("Process stopped:", message.Name)
	return nil
}
