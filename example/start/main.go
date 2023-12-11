package main

import (
	"errors"
	"log"

	"github.com/WaynePluto/go-lite"
)

func main() {
	l := lite.New()

	l.Use("/", func(ctx *lite.Context) {
		log.Printf("path: %v, method: %v", ctx.Path, ctx.Method)
		ctx.Next()
	})

	l.Use("/", func(ctx *lite.Context) {
		ctx.Next()
		if ctx.Err != nil {
			ctx.Json(500, ctx.Err.Error())
		}
	})

	l.GET("/", func(ctx *lite.Context) {
		ctx.JSON(ctx.Query())
	}, nil)

	l.GET("/err", func(ctx *lite.Context) {
		ctx.Err = errors.New("test get err")
	}, nil)

	l.GET("/ping/:id", func(ctx *lite.Context) {
		ctx.JSON(ctx.Params["id"])
	}, nil)

	l.GET("/headers", func(ctx *lite.Context) {
		ctx.JSON(ctx.Req.Header)
	}, nil)

	l.POST("/", func(ctx *lite.Context) {
		body, err := ctx.Body()
		if err != nil {
			return
		}
		ctx.JSON(body)
	})

	l.Run(":8000")
}
