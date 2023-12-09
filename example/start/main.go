package main

import (
	"github.com/WaynePluto/go-lite"
)

func main() {
	l := lite.New()

	l.Use("/ping", func(ctx *lite.Context) {
		ctx.Params["test"] = "test"
		ctx.Next()
	})

	l.Use("/ping/:id", func(ctx *lite.Context) {
		ctx.Params["test"] = ""
		ctx.Next()
	}, func(ctx *lite.Context) {
		ctx.Params["test"] = "test"
		ctx.Next()
	},
	)

	l.GET("/", func(ctx *lite.Context) {
		ctx.JSON("Hello,world")
	}, nil)
	l.GET("/ping/:id", func(ctx *lite.Context) {
		ctx.JSON(ctx.Params)
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
