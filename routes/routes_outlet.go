package routes

import (
	"net/url"
	"regexp"

	"ireul.com/neowx/types"
	"ireul.com/web"
)

// OutletWhiteList white list of url can be invoked
var OutletWhiteList = regexp.MustCompile(`\Ahttp(s)?:\/\/(\w+\.)?api\.weixin\.qq\.com\/`)

// OutletAction GET/POST to /outlet/:name
func OutletAction(ctx *web.Context, a types.Account) {
	url, err := url.QueryUnescape(ctx.Req.URL.RawQuery)
	if err != nil {
		ctx.Error(400, "url is malformatted: "+err.Error())
		return
	}
	if !OutletWhiteList.MatchString(url) {
		ctx.Error(400, "url is not permitted, make sure it's a Wechat API url")
		return
	}
	ctx.PlainText(200, []byte(a.AppID))
}
