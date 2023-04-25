package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func UploadsFileController(c *gin.Context) {
	file, err := c.FormFile("raw")
	if err != nil {
		log.Fatalln(err)
	}
	//	查找可执行文件目录
	exe, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	//exeDir := filepath.Dir(exe)
	exeDir := path.Dir(exe)

	fileName := uuid.New().String()
	uploads := filepath.Join(exeDir, "uploads")
	//	创建uploads
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	fullpath := path.Join("uploads", fileName+filepath.Ext(file.Filename))

	err = c.SaveUploadedFile(file, filepath.Join(exeDir, fullpath))
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
}
