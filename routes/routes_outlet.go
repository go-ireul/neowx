package routes

import (
	"net/url"
	"regexp"
	"strings"

	"ireul.com/neowx/store"
	"ireul.com/neowx/types"
	"ireul.com/web"
)

// AccessTokenPlaceholder placeholder for AccessToken in url
const AccessTokenPlaceholder = "$ACCESS_TOKEN"

// OutletWhiteList white list of url can be invoked
var OutletWhiteList = regexp.MustCompile(`\Ahttp(s)?:\/\/(\w+\.)?api\.weixin\.qq\.com\/`)

// OutletAction GET/POST to /outlet/:name
func OutletAction(ctx *web.Context, a types.Account, sto *store.Store) {
	// extract url
	url, err := url.QueryUnescape(ctx.Req.URL.RawQuery)
	if err != nil {
		ctx.Error(400, "url is malformatted: "+err.Error())
		return
	}
	// check whitelist
	if !OutletWhiteList.MatchString(url) {
		ctx.Error(400, "url is not permitted, make sure it's a Wechat API url")
		return
	}
	// fill AccessToken
	if strings.Contains(url, AccessTokenPlaceholder) {
		at, err := sto.GetAccessToken(a.AppID)
		if err != nil {
			ctx.Error(500, "failed to fetch access_token")
			return
		}
		url = strings.Replace(url, AccessTokenPlaceholder, at, -1)
	}
	ctx.PlainText(200, []byte(a.AppID))
}
