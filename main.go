package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"os/signal"
)

func main() {
	go func() {
		//gin.SetMode()
		r := gin.Default()
		r.GET("/", func(context *gin.Context) {
			context.String(200, "hello")
		})
		r.Run(":8080")
	}()

	//打开chrome 窗口下`
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:8080/")
	cmd.Start()

	//监听退出序号

	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, os.Interrupt)
	//select 会一直轮询直到chanSignal 有值
	select {
	case <-chanSignal:
		cmd.Process.Kill()
	}

}
