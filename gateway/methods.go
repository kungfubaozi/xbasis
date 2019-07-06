package main

import (
	"fmt"
	"konekko.me/xbasis/commons/errstate"
	"net/http"
)

func (r *request) get() {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s", r.path, r.c.Request.RequestURI), nil)
	req.Form = r.c.Request.Form

	if err != nil {
		r.c.JSON(200, errstate.ErrRequest)
		return
	}

	r.route(req)
}

func (r *request) post() {

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s%s", r.path, r.c.Request.RequestURI), r.c.Request.Body)
	if err != nil {
		r.json(errstate.ErrRequest)
		return
	}

	req.Form = r.c.Request.Form
	req.MultipartForm = r.c.Request.MultipartForm
	req.PostForm = r.c.Request.PostForm

	if err != nil {
		r.json(errstate.ErrRequest)
		return
	}

	r.route(req)
}

func (r *request) put() {

}

func (r *request) patch() {

}

func (r *request) options() {

}

func (r *request) delete() {

}

func (r *request) call(method string) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s", r.path, r.c.Request.RequestURI), nil)
	if err != nil {
		r.json(errstate.ErrRequest)
		return
	}
	req.Form = r.c.Request.Form
	switch method {
	case "GET":
		break
	case "POST":
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
