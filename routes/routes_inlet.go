package routes

import (
	"log"

	"ireul.com/neowx/types"
	"ireul.com/web"
)

// InletGet GET to /inlet
func InletGet(ctx *web.Context) {
	ctx.PlainText(200, []byte(ctx.Query("echostr")))
}

// InletPost POST to /inlet
func InletPost(ctx *web.Context, m types.Message) {
	log.Printf("%v\n", m)
	ctx.PlainText(200, []byte("success"))
}
