package main

import (
	"bytes"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func exampleAuthFunc(ctx *fasthttp.RequestCtx) bool {
	if bytes.HasPrefix(ctx.Path(), []byte("/protect")) {
		return false
	}
	return true
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("hello")
}

func panicHandler(ctx *fasthttp.RequestCtx) {
	panic("test panic")
}

func main() {
	//logger := zaplog.NewNoCallerLogger(false)
	//mo := fasthttpmiddleware.NewNormalMiddlewareOnion(exampleAuthFunc, false,logger)
	//moWithoutAuth := fasthttpmiddleware.NewMiddlewareOnion(
	//	fasthttpmiddleware.NewLogMiddleware(logger, false),
	//	fasthttpmiddleware.NewRecoverMiddleware(logger),
	//	//fasthttpmiddleware.
	//)
	router := fasthttprouter.New()
	router.GET("/", requestHandler)
	//router.GET("/protect", mo.Apply(requestHandler))
	//router.GET("/panic", mo.Apply(panicHandler))
	//router.GET("/noAuth", moWithoutAuth.Apply(requestHandler))

	fasthttp.ListenAndServe(":8000", router.Handler)

}
