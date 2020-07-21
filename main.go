package main

import (
	"fmt"
	"gptc-sync/syncer"
	"net/http"

	"gptc-sync/pkg/setting"
	"gptc-sync/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go syncer.SyncBlock()
	s.ListenAndServe()
}
