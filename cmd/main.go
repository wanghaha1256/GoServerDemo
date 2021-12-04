package main

import (
	"GoServerDemo/internal/routers"
	"GoServerDemo/internal/util/settings"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatalln(s.ListenAndServe())
}
