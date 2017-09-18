package routes

import "ireul.com/web"

// Mount mount routes *web.Web
func Mount(w *web.Web) {
	w.Get("/inlet", InletGet)
	w.Post("/inlet", DecodeWechatXML(), InletPost)
}
