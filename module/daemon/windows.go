//go:build windows

package daemon

import (
	"os/exec"
	"syscall"
)

func setupDaemon(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,                         // 콘솔 창 숨기기
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP, // 독립된 프로세스 그룹 생성
	}
}
