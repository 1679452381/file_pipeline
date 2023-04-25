package server

import (
	"embed"
	"file_pipeline/define"
	"file_pipeline/server/service"
	"file_pipeline/server/ws"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dist/*

var FS embed.FS

func Run() {
	//gin.SetMode()
	r := gin.Default()
	//r.GET("/", func(context *gin.Context) {
	//	context.String(200, "hello")
	//})

	// fs.Sub会从f(embed.FS)中提取所有前缀为"frontend/dist"的文件,并返回一个新的FS,包含这棵子树。
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	r.StaticFS("/static", http.FS(staticFiles))

	api := r.Group("/api")
	// 发送文本消息
	api.POST("/v1/texts", service.TextsController)
	//获取局域网ip
	api.GET("/v1/addresses", service.AddressesController)
	//文件下载
	r.GET("/uploads/:path", service.DownloadsController)
	//生成二维码
	api.GET("/v1/qrcodes", service.QrcodesController)
	//上传文件
	api.POST("/v1/files", service.UploadsFileController)
	//webscoket 传输文件
	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", func(c *gin.Context) {
		ws.WebsocketServe(hub, c)
	})
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

	r.Run(":" + define.Port)
}
