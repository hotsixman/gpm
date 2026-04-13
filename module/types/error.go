package types

import (
	"fmt"
)

// Server
type InvalidMessage struct {
	JSON string
}

func (m InvalidMessage) Error() string {
	return fmt.Sprintf("Invalid message: %s", m.JSON)
}

type UndefinedProcessNameError struct{}

func (_ UndefinedProcessNameError) Error() string {
	return "'name' field is not defined."
}

// PM
/*
특정 이름의 프로세스가 없습니다.
*/
type NoProcessError struct {
	Name string
}

func (this NoProcessError) Error() string {
	return fmt.Sprintf("There is no process named \"%s\"", this.Name)
}

/*
프로세스가 실행중입니다.
*/
type ProcessRunningError struct {
	Name string
}

func (this ProcessRunningError) Error() string {
	return fmt.Sprintf("Process \"%s\" is running.", this.Name)
}

/*
프로세스가 실행중이지 않습니다.
*/
type ProcessNotRunningError struct {
	Name string
}

func (this ProcessNotRunningError) Error() string {
	return fmt.Sprintf("Process \"%s\" is not running.", this.Name)
}

// cli/startfrom

type NoNameError struct {
	JsonPath string
}

func (this NoNameError) Error() string {
	return fmt.Sprintf("No 'name' key in %s", this.JsonPath)
}

type NoRunError struct {
	JsonPath string
}

func (this NoRunError) Error() string {
	return fmt.Sprintf("No 'run' key in %s", this.JsonPath)
}

type InvalidArgsError struct {
	JsonPath string
}

func (this InvalidArgsError) Error() string {
	return fmt.Sprintf("'args' member is invalid in %s", this.JsonPath)
}
