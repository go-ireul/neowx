package types

import "testing"

func TestRuleMatches(t *testing.T) {
	r := Rule{
		Match: map[string]string{
			"Content":    "world",
			"CreateTime": `/^\d{2}$/`,
		},
	}
	rq := WxReq{
		Content:    "wOrld ",
		CreateTime: 333,
	}
	ok, err := r.Matches(rq)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Error("failed 1")
	}

	rq.Content = "na"

	ok, err = r.Matches(rq)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Error("failed 2")
	}

	rq.Content = "world"
	rq.CreateTime = 33

	ok, err = r.Matches(rq)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("failed 3")
	}
}
