package main

import (
	"github.com/gin-gonic/gin"
	"os/exec"
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
	//打开chrome 窗口
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:8080", "--window-size 400 400")
	cmd.Start()
	select {}
}
