package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func DownloadsController(c *gin.Context) {
	//获取文件路径
	filePath := c.Param("path")
	if filePath != "" {
		//将网络路径变成本地路径
		//查看当前可执行文件路径
		exe, err := os.Executable()
		if err != nil {
			log.Fatalln(err)
		}
		dir := path.Dir(exe)
		//读取本地文件，写到http相应
		target := filepath.Join(dir, filePath)
		fmt.Println(target)

		c.Header("Content-Disposition", "attachment; filename="+filePath)
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(target)

	} else {
		c.Status(http.StatusBadRequest)
	}
}
