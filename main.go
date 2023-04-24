package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

//embed.FS:用于嵌入文件,如通过go:embed嵌入本地文件

//go:embed frontend/dist/*
var FS embed.FS

func main() {
	go func() {
		//gin.SetMode()
		r := gin.Default()
		//r.GET("/", func(context *gin.Context) {
		//	context.String(200, "hello")
		//})

		// fs.Sub会从f(embed.FS)中提取所有前缀为"frontend/dist"的文件,并返回一个新的FS,包含这棵子树。
		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		r.StaticFS("/static", http.FS(staticFiles))
		// 静态文件不存在 渲染html
		r.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			if strings.HasPrefix(path, "/static/") {
				reader, err := staticFiles.Open("index.html")
				if err != nil {
					log.Fatalln(err)
				}
				defer reader.Close()
				stat, err := reader.Stat()
				if err != nil {
					log.Fatalln(err)
				}
				c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
			} else {
				c.Status(http.StatusNotFound)
			}
		})

		r.Run(":8080")
	}()

	//打开chrome 窗口下`
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:8080/static/index.html")
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
