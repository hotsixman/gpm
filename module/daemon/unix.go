//go:build !windows

package daemon

import (
	"os/exec"
	"syscall"
)

func setupDaemon(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true, // 새로운 세션 시작 (제어 터미널 분리)
	}
}
