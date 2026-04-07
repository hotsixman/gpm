package pm

import (
	"gpm/module/types"
)

func (pm *PM) Start(startMessage types.StartMessage) error {
	formerProcess := pm.process[startMessage.Name]
	if formerProcess == nil {
		return pm.initProcess(startMessage, nil)
	} else {
		return &types.ProcessRunningError{Name: startMessage.Name}
	}
}
