package main

import (
	"GoServerDemo/configs"
	"GoServerDemo/internal/routers"
	"GoServerDemo/internal/util/controllers"
	"GoServerDemo/internal/util/settings"
	"GoServerDemo/pkg"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
