package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func TextsController(c *gin.Context) {
	var json struct {
		Raw string `json:"raw"`
	}
	//	获取输入数据
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	} else {
		//  获取go可执行文件目录
		exe, err := os.Executable()
		if err != nil {
			log.Fatalln(err)
		}

		//	在该目录创建 uploads目录
		//  要使用filepath获取文件路径
		dir := filepath.Dir(exe)
		uploads := filepath.Join(dir, "uploads")
		err = os.MkdirAll(uploads, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		//  将文本保存到文件中
		fileName := uuid.New().String()
		fullPath := filepath.Join("uploads", fileName+".txt")
		err = os.WriteFile(filepath.Join(dir, fullPath), []byte(json.Raw), 0644)
		if err != nil {
			log.Fatalln(err)
		}
		//  返回该文件的下载路径
		c.JSON(http.StatusOK, gin.H{"url": "/" + fullPath})
	}
}
