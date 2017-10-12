package routes

import (
	"encoding/base64"
	"errors"
	"log"
	"net/url"
	"strings"

	"ireul.com/neowx/store"
	"ireul.com/neowx/types"
	"ireul.com/neowx/wx"
	"ireul.com/web"
)

// OAuthAuthorizeAction action for OAuth relay
func OAuthAuthorizeAction(ctx *web.Context, a types.Account, sto *store.Store, cfg types.Config) {
	// extract url
	link, err := url.QueryUnescape(ctx.Req.URL.RawQuery)
	if err != nil {
		ctx.Error(400, "url is malformatted: "+err.Error())
		return
	}
	if len(link) == 0 {
		link = ctx.Header().Get("Referer")
	}
	if len(link) == 0 {
		ctx.Error(400, "target url not set")
		return
	}
	// validate url
	if err = validateOAuthURL(link, cfg.OAuthDomains); err != nil {
		ctx.Error(400, err.Error())
		return
	}
	// cache url
	next := base64.RawURLEncoding.EncodeToString([]byte(link))
	// build wechat oauth url
	wxURL := wx.BuildOAuthURL(wx.OAuthOptions{
		AppID:       a.AppID,
		CallbackURL: cfg.URL + "/oauth/" + a.Name + "/callback/" + next,
		Scope:       wx.ScopeUserInfo,
	})
	ctx.Redirect(wxURL)
}

// OAuthCallbackAction action for OAuth relay
func OAuthCallbackAction(ctx *web.Context, a types.Account, sto *store.Store, cfg types.Config) {
	code := ctx.Query("code")
	next := ctx.Params(":next")

	// extract url
	rlink, err := base64.RawURLEncoding.DecodeString(next)
	if err != nil || len(rlink) == 0 {
		ctx.Error(400, "failed to decode url: "+err.Error())
		return
	}
	link := string(rlink)

	// get user token
	resp, err := wx.GetUserToken(wx.GetUserTokenOptions{
		AppID:     a.AppID,
		AppSecret: a.AppSecret,
		Code:      code,
	})
	if err != nil {
		ctx.Error(400, "failed to get user token: "+err.Error())
		return
	}
	log.Println(link)
	log.Println(resp.OpenID)
	ctx.PlainText(200, []byte("hello"))
}

func validateOAuthURL(link string, ds []string) (err error) {
	var u *url.URL
	if u, err = url.Parse(link); err != nil {
		err = errors.New("url is malformatted: " + err.Error())
		return
	}
	if len(ds) > 0 {
		for _, v := range ds {
			if strings.ToLower(strings.TrimSpace(v)) == strings.ToLower(strings.TrimSpace(u.Host)) {
				return nil
			}
		}
		err = errors.New("url is not permitted")
	}
	return nil
}
