package router

import (
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"gitlab.com/tokend/hgate/internal/helpers"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/tokend/hgate/internal/handlers"
	"gitlab.com/tokend/keypair"
)

func NewRouter(signer keypair.Full, horizonURL *url.URL, logger *logan.Entry) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(logger),
		middleware.URLFormat,
		middleware.SetHeader("Content-Type", "application/json"),
	)

	submitTx, err := helpers.GetTxSubmitter(signer, horizonURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tx submitter")
	}
	signRequest := helpers.GetRequestSigner(signer)
	proxy := helpers.GetProxy(horizonURL)

	r.Patch("/assets/{code}", handlers.UpdateAsset(submitTx, logger))
	r.Post("/create_kyc_request", handlers.CreateKYCRequest(submitTx, logger))

	// Proxy all other requests to horizon
	r.Handle("/*", http.HandlerFunc(handlers.RedirectHandler(signRequest, proxy, logger)))

	return r, nil
}
