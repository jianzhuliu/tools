/*
浏览器相关工具
*/

package browser

import (
	"os/exec"
	"runtime"
	"time"

	"fmt"
)

//浏览器自动打开 url
//return bool
func OpenUrl(targetUrl string) bool {
	args := []string{}

	switch runtime.GOOS {
	//操作系统判断，目前只支持 windows
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		return false
	}

	//构造命令
	cmd := exec.Command(args[0], append(args[1:], targetUrl)...)

	return CommandExecCheck(cmd, 3)
}

//判断命令是否执行且未超时
//return bool
func CommandExecCheck(cmd *exec.Cmd, t time.Duration) bool {
	if cmd.Start() != nil {
		return false
	}

	//另外开启 goroutine 来存储执行后的错误信息
	errChan := make(chan error, 1)
	go func() {
		errChan <- cmd.Wait()
	}()

	//多路复用
	select {
	case <-time.After(t):
		//超时处理
		return false
	case err := <-errChan:
		return err == nil
	}
}

//用于开启一个新线程，带有提示输出
func OpenWithNotice(addr string) {
	targetUrl := fmt.Sprintf("http://%s", addr)
	if OpenUrl(targetUrl) {
		fmt.Printf("your web browser is open. if not, please open yourself, and visit %s \n", targetUrl)
	} else {
		fmt.Printf("please open your web browser and visit %s \n", targetUrl)
	}
}
