package server

import (
	"gitlab.com/tokend/hgate/server/problem"
	"encoding/json"
	"gitlab.com/distributed_lab/logan"
	"net/http"
)

const MimeJSON = "application/json"

type HandlerI interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	HandleRequest(w http.ResponseWriter, r *http.Request)
	Render(w http.ResponseWriter, data interface{})
	FinishRequest()
}

type Handler struct {
	Log         *logan.Entry
	Err         error
	ContentType string
	Method      string
	rawBody     []byte
	R           *http.Request
	W           http.ResponseWriter
}

func (action *Handler) FinishRequest() {
	action.Log.WithField("path", action.R.URL.Path).Info("Finished request")
	action.Err = nil
	action.R = nil
	action.W = nil
	action.rawBody = nil
}

func (action *Handler) SetRequest(r *http.Request) {
	action.R = r
}

func (action *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (action *Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (action *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (action *Handler) Render(w http.ResponseWriter, data interface{}) {
	if action.Err != nil {
		problem.Render(w, action.Err)
		return
	}

	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
