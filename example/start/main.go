package main

import (
	"log"

	"github.com/WaynePluto/go-lite"
)

func main() {
	l := lite.New()

	l.Use("/", func(ctx *lite.Context) {
		log.Printf("path: %v, method: %v", ctx.Path, ctx.Method)
		ctx.Next()
	})

	l.GET("/", func(ctx *lite.Context) {
		ctx.JSON(nil)
	}, nil)

	l.GET("/ping/:id", func(ctx *lite.Context) {
		ctx.JSON(ctx.Params["id"])
	}, nil)

	l.GET("/headers", func(ctx *lite.Context) {
		ctx.JSON(ctx.Req.Header)
	}, nil)

	l.POST("/", func(ctx *lite.Context) {
		body, err := ctx.GetReqBody()
		if err != nil {
			return
		}
		ctx.JSON(body)
	})

	l.Run(":8000")
}
