package handlers

import (
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func CreateAssetCreationRequest(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewCreateAssetCreationRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	env, err := buildAndSign(r, &xdrbuild.CreateAsset{
		Code:                     request.Code,
		MaxIssuanceAmount:        uint64(request.MaxIssuanceAmount),
		PreIssuanceSigner:        request.PreIssuanceAssetSigner,
		InitialPreIssuanceAmount: uint64(request.InitialPreissuedAmount),
		TrailingDigitsCount:      request.TrailingDigitsCount,
		Policies:                 request.Policies,
		Type:                     request.Type,
		CreatorDetails:           request.CreatorDetails,
		AllTasks:                 request.AllTasks,
	})

	if err != nil {
		log.WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}
