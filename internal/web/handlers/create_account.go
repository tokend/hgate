package handlers

import (
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewCreateAccountRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	signersData := make([]xdrbuild.SignerData, 0, len(request.Signers))

	for _, s := range request.Signers {
		signersData = append(signersData, xdrbuild.SignerData{
			PublicKey: s.PublicKey,
			RoleID:    s.SignerRole,
			Weight:    s.Weight,
			Identity:  s.Identity,
			Details:   s.Details,
		})
	}

	env, err := buildAndSign(r, &xdrbuild.CreateAccount{
		Destination: request.Destination,
		Referrer:    request.Referrer,
		RoleID:      request.AccountRole,
		Signers:     signersData,
	})

	if err != nil {
		log.WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}
