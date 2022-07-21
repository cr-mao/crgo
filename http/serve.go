package http

import (
	"net/http"
	"time"

	"crgo/http/routers"
	"crgo/infra/conf"
	"crgo/infra/log"
)

func NewServe() *http.Server {
	router := routers.NewRouter()

	addr := conf.GetString("http_addr") + ":" + conf.GetString("http_port")
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Infof("http listening on %s", addr)
	return s
}
