package service

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func QrcodesController(c *gin.Context) {
	content := c.Query("content")
	if content != "" {
		//	生成二维码
		pngFile, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatalln(err)
		}
		c.Data(http.StatusOK, "image/png", pngFile)
	}
	c.Status(http.StatusBadRequest)
}
