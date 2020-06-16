package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"myweb/handler"
	"net/http"
)

func main() {
	server := gin.New()

	server.GET("/index", handler.Index)

	if err := http.ListenAndServe(":19020", server); err != nil {
		log.Fatalf("server return err:%v", err)
	}
}
