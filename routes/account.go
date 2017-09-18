package routes

import (
	"ireul.com/neowx/types"
	"ireul.com/web"
)

// AccountFilter middleware retrieve account by name
func AccountFilter(k string) interface{} {
	return func(ctx *web.Context, cfg types.Config) {
		name := ctx.Params(":name")
		a := cfg.AccountByName(name)
		if len(a.Name) == 0 {
			a = cfg.DefaultAccount()
			if len(a.Name) == 0 {
				ctx.Error(400, "Wechat MP account not found: "+name)
				return
			}
		}
		ctx.Map(a)
	}
}
