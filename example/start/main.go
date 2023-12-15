package main

import (
	"errors"
	"log"

	"github.com/WaynePluto/go-lite"
	"github.com/WaynePluto/go-lite/middlewares"
)

func main() {
	l := lite.New()

	l.Use("/", func(ctx *lite.Context) {
		log.Printf("path: %v, method: %v", ctx.Path, ctx.Method)
		ctx.Next()
	})

	l.Use("/", middlewares.Error(nil))

	l.GET("/", func(ctx *lite.Context) {
		ctx.JSON(ctx.Query())
	}, nil)

	// 测试错误处理中间件
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
		body, ok := ctx.Body()
		if ok {
			ctx.JSON(body)
		}
	}, nil)

	l.Run(":8000")
}
