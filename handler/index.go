package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Resp struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func Index(ctx *gin.Context) {
	log.Printf("Index Req:%v", ctx.Request)
	ctx.JSON(200, &Resp{
		Name: "Donald Trump",
		Age:  88,
	})
}
