package xbasisgateway

import (
	"fmt"
	"konekko.me/xbasis/commons/errstate"
	"net/http"
)

func (r *request) call(method string) {
	req, err := http.NewRequest(method, fmt.Sprintf("http://%s%s", r.path, r.c.Request.RequestURI), r.c.Request.Body)
	r.requestMethod = method
	r.requestPath = r.c.Request.RequestURI
	if err != nil {
		r.json(errstate.ErrRequest)
		return
	}
	req.Form = r.c.Request.Form
	switch method {
	case "GET":
		break
	case "POST":
		req.MultipartForm = r.c.Request.MultipartForm
		req.PostForm = r.c.Request.PostForm
		break
	case "PUT":
		break
	case "PATCH":
		break
	case "HEAD":
		break
	case "OPTIONS":
		break
	case "DELETE":
		break
	case "TRACE":
		break
	case "CONNECT":
		break
	default:
		r.json(errstate.ErrRequest)
		return
	}
	r.route(req)
}
