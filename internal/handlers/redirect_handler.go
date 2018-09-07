package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/hgate/internal/helpers"
)

func RedirectHandler(signRequest helpers.RequestSigner, proxy helpers.Proxy, logger *logan.Entry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := signRequest(r)
		if err != nil {
			logger.WithError(err).Error("Failed to sign request.")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		proxy(w, r)
	}
}
