package main

import (
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//var ui lorca.UI
	//--disable-sync 取消同步功能
	//--disable-translate 取消翻译功能
	ui, err := lorca.New("https://www.baidu.com", "", 800, 600, "--disable-sync", "--disable-translate")
	if err != nil {
		log.Fatalln(err.Error())
	}
	// 监听 中断信号
	chSignal := make(chan os.Signal, 1)
	// Notify 订阅 监听到中断信号
	//syscall.SIGINT 中断信号
	//syscall.SIGTERM 终止信号
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ui.Done():
	case <-chSignal:
	}

	ui.Close()
}
