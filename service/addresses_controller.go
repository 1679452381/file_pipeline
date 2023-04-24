package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

func AddressesController(c *gin.Context) {
	// 获取当前电脑 多有ip
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln(err)
	}
	result := make([]string, 0)

	for _, addr := range addrs {
		// !ipNet.IP.IsLoopback 是网卡并且不是本地环回网卡
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				result = append(result, ipv4.String())
			}

		}
	}
	c.JSON(http.StatusOK, gin.H{"addresses": result})
}
