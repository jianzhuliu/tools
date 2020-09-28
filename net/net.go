/*
网络相关
*/

package net

import (
	"fmt"
	"os/exec"
	"runtime"
)

//检查端口是否被占用
//return bool | true 被占用
func CheckPortInUse(port int) bool {
	var cmdName string
	var args []string
	var cmdStr string

	switch runtime.GOOS {
	case "windows":
		// cmd /c netstat -ano | findstr 1024
		cmdName = "cmd"
		cmdStr = fmt.Sprintf("netstat -ano | findstr %d", port)
		args = []string{"/c", cmdStr}
	case "linux":
		// sh -c lsof -i:1024
		cmdName = "sh"
		cmdStr = fmt.Sprintf("lsof -i:%d", port)
		args = []string{"-c", cmdStr}
	default:
		return false
	}

	cmd := exec.Command(cmdName, args...)
	output, _ := cmd.Output() //只关注结果，异常情况忽略
	if len(output) > 0 {
		return true
	}

	return false
}
