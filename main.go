package main

import (
	"file_pipeline/define"
	"file_pipeline/server"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
)

//embed.FS:用于嵌入文件,如通过go:embed嵌入本地文件

func main() {
	go func() {
		server.Run()
	}()

	chanChrome := make(chan struct{})
	//cmd打开chrome窗口
	go CmdStartChrome(chanChrome)
	fmt.Println("关闭了")

	//监听退出序号
	chanSignal := InterruptSignal()
	//select 会一直轮询直到chanSignal 有值
	select {
	case <-chanSignal:
	case <-chanChrome:
		os.Exit(0)
		return
	}
}

// cmd
func CmdStartChrome(chanChrome chan struct{}) *exec.Cmd {
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:"+define.Port+"/static/index.html")
	cmd.Start()
	cmd.Wait()
	chanChrome <- struct{}{}
	return cmd
}

// 监听退出信号
func InterruptSignal() chan os.Signal {
	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, os.Interrupt)
	return chanSignal
}
