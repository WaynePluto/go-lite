package main

import (
	"log"

	"github.com/WaynePluto/go-lite"
	"github.com/WaynePluto/go-lite/utils"
)

type userQuery struct {
	Page     int "json:page"
	PageSize int "json:pageSize"
}

type userBody struct {
	Phone string `json:"phone" type:"string"`
	Pass  string `json:"pass" type:"string"`
}

// 定义api参数
var apiParam = lite.ApiParam{
	Query: utils.StructToSlice(userQuery{}),
	Body:  utils.StructToMap(userBody{}),
}

// 定义控制器
func controller(ctx *lite.Context) {

}

func main() {
	l := lite.New()

	l.Use("/", func(ctx *lite.Context) {
		log.Printf("path: %v, method: %v", ctx.Path, ctx.Method)
		ctx.Next()
	})

	l.POST("/user/:id/books", controller, &apiParam)

}
