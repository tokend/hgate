package server

import (
	"gitlab.com/tokend/hgate/server/problem"
	"net/http"
)

type Mux struct {
	*http.ServeMux
}

func NewMux() *Mux {
	return &Mux{ServeMux: http.NewServeMux()}
}

func (mux *Mux) Get(path string, actionInit func(r *http.Request) HandlerI) {
	mux.handle(http.MethodGet, path, actionInit)
}

func (mux *Mux) Post(path string, actionInit func(r *http.Request) HandlerI) {
	mux.handle(http.MethodPost, path, actionInit)
}

func (mux *Mux) handle(method, path string, actionInit func(r *http.Request) HandlerI) {
	hf := func(w http.ResponseWriter, r *http.Request) {
		if method != "" && r.Method != method {
			problem.Render(w, problem.NotAcceptable)
			return
		}
		action := actionInit(r)
		switch method {
		case http.MethodGet:
			action.Get(w, r)
			break
		case http.MethodPost:
			action.Post(w, r)
			break
		default:
			action.HandleRequest(w, r)
		}
		action.FinishRequest()
	}
	mux.HandleFunc(path, hf)
}
