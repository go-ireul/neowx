package routes

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"ireul.com/com"
	"ireul.com/neowx/types"
	"ireul.com/neowx/wx"
	"ireul.com/web"
)

// InletGet GET to /inlet
func InletGet(ctx *web.Context) {
	ctx.PlainText(200, []byte(ctx.Query("echostr")))
}

// InletPost POST to /inlet
func InletPost(ctx *web.Context, m wx.Message, cfg types.Config) {
	for _, rule := range cfg.Rules {
		// match
		ok, err := rule.Matches(m)
		// if failed to execute match, basically regexp is wrongly written, mute and return
		if err != nil {
			log.Printf("--- Failed to execute match:\n%s\n%v\n--- \n", err.Error(), m)
			ctx.PlainText(200, []byte("success"))
			return
		}
		// if matched
		if ok {
			// if HTTPSync, make request, pipe and return
			if len(rule.HTTPSync) > 0 {
				// execute POST with Message marshalled as JSON
				resp, err := relayMessage(rule.HTTPSync, m)
				// if failed, mute
				if err != nil {
					log.Printf("--- Failed to POST HTTPSync to %s\n%s\n--- \n", rule.HTTPSync, err.Error())
					ctx.PlainText(200, []byte("success"))
					return
				}
				// pipe and return
				if err = com.HTTPPipeResponse(ctx.Resp, resp); err != nil {
					log.Printf("--- Failed to pipe HTTPSync to %s\n%s\n--- \n", rule.HTTPSync, err.Error())
				}
				return
			}

			// if HTTPAsync, go async request
			if len(rule.HTTPAsync) > 0 {
				go relayMessage(rule.HTTPAsync, m)
			}
			// write text or mute
			if len(rule.Text) > 0 {
				resp := wx.TextReply{
					FromUserName: com.NewCDATA(m.ToUserName),
					ToUserName:   com.NewCDATA(m.FromUserName),
					CreateTime:   com.ToStr(time.Now().Unix()),
					MsgType:      com.NewCDATA(wx.Text),
					Content:      com.NewCDATA(rule.Text),
				}
				ctx.XML(200, resp)
			} else {
				ctx.PlainText(200, []byte("success"))
			}

			return
		}
		// continue to next rule
	}

	// if no rule matches found, mute
	ctx.PlainText(200, []byte("success"))
}

func relayMessage(url string, m wx.Message) (*http.Response, error) {
	// marshal to json
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	rdr := bytes.NewReader(b)
	// post as application/json
	return http.Post(url, "application/json", rdr)
}
