package routes

import (
	"net/url"
	"regexp"

	"ireul.com/neowx/store"
	"ireul.com/neowx/types"
	"ireul.com/web"
)

// OAuthWhiteList whitelist for OAuth source page
var OAuthWhiteList = regexp.MustCompile(`\Ahttp(s)?:\/\/.+\.pagoda.com.cn\/`)

// OAuthAction action for OAuth relay
func OAuthAction(ctx *web.Context, a types.Account, sto *store.Store) {
	// extract url
	ul, err := url.QueryUnescape(ctx.Req.URL.RawQuery)
	if err != nil {
		ctx.Error(400, "url is malformatted: "+err.Error())
		return
	}
	if len(ul) == 0 {
		ul = ctx.Header().Get("Referer")
	}
	if len(ul) == 0 {
		ctx.Error(400, "target url not set")
		return
	}
	ctx.PlainText(200, []byte(ul))
}
