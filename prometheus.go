package common

import (
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	go func ()  {
		err := http.ListenAndServe("0.0.0.0:" + strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("启动普罗米修斯失败" + err.Error())
		}
		log.Println("启动普罗米修斯成功，端口：", strconv.Itoa(port))
	}()
}