package wx

import (
	"errors"
	"net/http"

	"ireul.com/com"
)

// Resp abstract a API response with errcode and errmsg
type Resp interface {
	GetErrorCode() int
	GetErrorMessage() string
}

// BaseResp basic resp for wx api
type BaseResp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

// GetErrorCode implements interface Resp
func (br BaseResp) GetErrorCode() int {
	return br.ErrorCode
}

// GetErrorMessage implements interface Resp
func (br BaseResp) GetErrorMessage() string {
	return br.ErrorMessage
}

func combineRespError(r Resp, err error) error {
	if err != nil {
		return err
	}
	if r.GetErrorCode() != 0 {
		return errors.New(r.GetErrorMessage())
	}
	return nil
}

// GetAPI make a HTTP Get request to Wechat API
func GetAPI(link string, out interface{}) error {
	err := com.HttpGetJSON(http.DefaultClient, link, out)
	if k, ok := out.(Resp); ok {
		return combineRespError(k, err)
	}
	return err
}
