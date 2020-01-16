package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := getKeys(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get keys")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	err = signRequest(r, keys.Signer)
	if err != nil {
		Log(r).WithError(err).Error("failed to sign request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	Proxy(w, r)
}
