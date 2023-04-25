package main

import (
	"file_pipeline/define"
	"file_pipeline/server"
	"os"
	"os/exec"
	"os/signal"
)

//embed.FS:用于嵌入文件,如通过go:embed嵌入本地文件

func main() {
	go func() {
		server.Run()
	}()

	//打开chrome 窗口下
	cmd := Cmd()
	cmd.Start()

	//监听退出序号
	chanSignal := InterruptSignal()
	//select 会一直轮询直到chanSignal 有值
	select {
	case <-chanSignal:
		cmd.Process.Kill()
	}
}

// cmd
func Cmd() *exec.Cmd {
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	return exec.Command(chromePath, "--app=http://localhost:"+define.Port+"/static/index.html")
}

// 监听退出序号
func InterruptSignal() chan os.Signal {
	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, os.Interrupt)
	return chanSignal
}
