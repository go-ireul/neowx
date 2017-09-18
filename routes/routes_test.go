package routes

import "testing"

func TestOutletWhiteList(t *testing.T) {
	urls := []string{
		"http://api.weixin.qq.com/dfafa",
		"https://sh.api.weixin.qq.com/dfasfas",
		"https://sh.api.weixin.qq.com/dasfdas",
		"http://hk.api.weixin.qq.com/somepath/somewhere",
	}
	for _, url := range urls {
		if !OutletWhiteList.MatchString(url) {
			t.Error("failed to allow url: " + url)
		}
	}

	urls = []string{
		"internal://api.weixin.qq.com/dfafa",
		"https://ss.dash.api.weixin.qq.com/dfasfas",
		"https://sh.api.weixin.qq-online.com/dasfdas",
		"http://weixin.qq.com/somepath/somewhere",
	}
	for _, url := range urls {
		if OutletWhiteList.MatchString(url) {
			t.Error("failed to deny url: " + url)
		}
	}
}
