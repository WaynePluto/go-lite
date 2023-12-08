package main

import (
	lite "lite/core"
)

func main() {
	l := lite.New()
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