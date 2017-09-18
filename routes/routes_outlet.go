package routes

import (
	"ireul.com/neowx/types"
	"ireul.com/web"
)

// OutletGet GET to /outlet/:name
func OutletGet(ctx *web.Context, a types.Account) {
	ctx.PlainText(200, []byte(a.AppID))
}

// OutletPost POST to /outlet/:name
func OutletPost(ctx *web.Context, a types.Account) {
	ctx.PlainText(200, []byte(a.AppID))
}
