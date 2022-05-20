package httpproject

import (
	"context"
	"crgo/infra/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"crgo/httpproject/routers"
	"crgo/infra/conf"
)

func Run() error {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           conf.Net.HTTP_ADDR + ":" + conf.Net.HTTP_PORT,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	//更新黑名单 内存
	go func() {
		util.WatchBlacklist()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}
