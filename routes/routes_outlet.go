package routes

import (
	"errors"
	"net/url"
	"strings"

	"ireul.com/neowx/store"
	"ireul.com/neowx/types"
	"ireul.com/web"
)

// AccessTokenPlaceholder placeholder for AccessToken in url
const AccessTokenPlaceholder = "$ACCESS_TOKEN"

// OutletAction GET/POST to /outlet/:name
func OutletAction(ctx *web.Context, a types.Account, sto *store.Store) {
	// extract url
	link, err := url.QueryUnescape(ctx.Req.URL.RawQuery)
	if err != nil {
		ctx.Error(400, "url is malformatted: "+err.Error())
		return
	}
	// check whitelist
	if err = validateOutletURL(link); err != nil {
		ctx.Error(400, err.Error())
		return
	}
	// fill AccessToken
	if strings.Contains(link, AccessTokenPlaceholder) {
		at, err := sto.GetAccessToken(a.AppID)
		if err != nil {
			ctx.Error(500, "failed to fetch access_token")
			return
		}
		link = strings.Replace(link, AccessTokenPlaceholder, at, -1)
	}
	ctx.PlainText(200, []byte(a.AppID))
}

func validateOutletURL(link string) (err error) {
	var u *url.URL
	if u, err = url.Parse(link); err != nil {
		err = errors.New("url is malformatted: " + err.Error())
		return
	}
	if !strings.HasSuffix(strings.ToLower(u.Host), "api.weixin.qq.com") {
		err = errors.New("url is not permitted, make sure it's a Wechat API url")
		return
	}
	return nil
}
