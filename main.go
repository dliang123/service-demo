package main

import (
	"github.com/go-kit/kit/log"
	"fmt"
	"os"
	"service-demo/service"
	"net/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	fmt.Print("hello world")
	logger.Log("nice")
	svc := &service.ChinaHelloService{}
	handle := service.NewRounter(svc)
	err := http.ListenAndServe(":8080", handle)
	if err != nil {
		logger.Log("ListenAndServe err: ", err)
	}
}
