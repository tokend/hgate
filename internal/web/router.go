package web

import (
	"github.com/go-chi/chi"
	"github.com/tokend/hgate/internal/config"
	"github.com/tokend/hgate/internal/helpers"
	"github.com/tokend/hgate/internal/web/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func Router(cfg config.Config, builder *xdrbuild.Builder) chi.Router {
	r := chi.NewRouter()
	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			handlers.CtxLog(cfg.Log()),
			handlers.CtxSubmitter(cfg.Submit()),
			handlers.CtxBuilder(builder),
			handlers.CtxKeys(cfg.Keys()),
			handlers.CtxProxy(helpers.GetProxy(cfg.HorizonURL())),
		),
	)

	r.Handle("/*", http.HandlerFunc(handlers.RedirectHandler))

	r.Route("/integrations/hgate", func(r chi.Router) {
		r.Post("/create_account", handlers.CreateAccount)
		r.Post("/change_role", handlers.ChangeRole)

		r.Post("/payment", handlers.Payment)

		r.Post("/create_asset", handlers.CreateAsset)
		r.Post("/update_asset", handlers.UpdateAsset)

		r.Post("/manage_signer", handlers.ManageSigner)

		r.Post("/keypair", handlers.GenerateKeypair)
	})

	return r
}
