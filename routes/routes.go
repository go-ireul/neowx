package routes

import "ireul.com/web"

// Mount mount routes *web.Web
func Mount(w *web.Web) {
	w.Get("/inlet", InletGet)
	w.Post("/inlet", DecodeXMLFilter(), InletPost)
	w.Get("/outlet/:name", AccountFilter(":name"), OutletAction)
	w.Get("/oauth/:name", AccountFilter(":name"), OAuthAction)
}
