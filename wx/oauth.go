package wx

import (
	"net/url"
)

// OAuthOptions options to build oauth url
type OAuthOptions struct {
	AppID       string
	CallbackURL string
	State       string
	Scope       string
}

const (
	// ScopeBase scope without prompt to authorize
	ScopeBase = "snsapi_base"
	// ScopeUserInfo scope full user info
	ScopeUserInfo = "snsapi_userinfo"
)

// BuildOAuthURL build a OAuth url for Wechat API
func BuildOAuthURL(opt OAuthOptions) string {
	if len(opt.Scope) == 0 {
		opt.Scope = ScopeBase
	}
	opt.CallbackURL = url.QueryEscape(opt.CallbackURL)
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + opt.AppID + "&redirect_uri=" + opt.CallbackURL + "&response_type=code&scope=snsapi_userinfo&state=" + opt.State + "#wechat_redirect"
}

// GetUserTokenOptions get personal access_token from Wechat API
type GetUserTokenOptions struct {
	AppID     string
	AppSecret string
	Code      string
}

// UserTokenResp access token response
type UserTokenResp struct {
	BaseResp
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// BuildGetUserTokenURL build GetUserTokenURL
func BuildGetUserTokenURL(opt GetUserTokenOptions) string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + opt.AppID + "&secret=" + opt.AppSecret + "&code=" + opt.Code + "&grant_type=authorization_code"
}

// GetUserToken fetch access token from Wechat API
func GetUserToken(opt GetUserTokenOptions) (UserTokenResp, error) {
	link := BuildGetUserTokenURL(opt)
	resp := UserTokenResp{}
	return resp, GetAPI(link, &resp)
}
