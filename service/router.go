package service

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func NewRounter(svr Service) http.Handler {

	router := httprouter.New()
	router.Handler("POST", "/hello/sayHello", NewSayHelloHandler(svr))

	return router

}
