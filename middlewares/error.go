package middlewares

import "github.com/WaynePluto/go-lite"

// 错误处理中间件
func Error(fn func(error)) func(ctx *lite.Context) {
	return func(ctx *lite.Context) {
		ctx.Next()
		if ctx.Err != nil {
			if fn != nil {
				fn(ctx.Err)
			}
			ctx.Json(500, ctx.Err.Error())
		}
	}
}
