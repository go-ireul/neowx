package routes

import (
	"ireul.com/web"
)

// OutletGet GET to /outlet
func OutletGet(ctx *web.Context) {
	name := ctx.Params(":name")
	ctx.PlainText(200, []byte(name))
}

// OutletPost POST to /outlet
func OutletPost(ctx *web.Context) {
	name := ctx.Params(":name")
	ctx.PlainText(200, []byte(name))
}
