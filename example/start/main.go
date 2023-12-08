package main

import "github.com/WaynePluto/go-lite"

func main() {
	l := lite.New()

	l.Use("/ping/:id", func(ctx *lite.Context) {
		ctx.Params["test"] = "test"
	})

	l.GET("/", func(c *lite.Context) {
		c.JSON("Hello,world")
	})
	l.GET("/ping/:id", func(c *lite.Context) {
		c.JSON(c.Params)
	})
	l.GET("/headers", func(c *lite.Context) {
		c.JSON(c.Req.Header)
	})

	l.POST("/", func(ctx *lite.Context) {
		body, err := ctx.GetReqBody()

		if err != nil {
			return
		}
		ctx.JSON(body)
	})

	l.Run(":8000")
}
