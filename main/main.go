/*
	1) "Hello World for Macaron"

	func main() {
    m := macaron.Classic()
    m.Get("/", func() string {
        return "Hello world!"
    })
    m.Run()

	2) object in Macaron
	*macaron.Context - HTTP context
	*log.Logger - Macaron global log
	http.ResponseWriter - HTTP response writer
	*http.Request - HTTP request obj

	3) middleware
	m.Use(func() {
  	// 处理中间件事务
	})

	4) Handlers
	m.Handlers(
    	Middleware1,
    	Middleware2,
    	Middleware3,
	)

	5) Hello world in Middleware
	m.Use(func(ctx *macaron.Context) {
    if ctx.Req.Header.Get("X-API-KEY") != "secret123" {
        ctx.Resp.WriteHeader(http.StatusUnauthorized)
    }
})
}


*/
package main

import (
	"runtime"
	"tcf/web"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	// init a Macaron object
	web.RunWeb()
}
