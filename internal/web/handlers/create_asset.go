package handlers

import (
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewCreateAssetRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	requestID := uint64(0)
	if request.RequestId != nil {
		requestID = *request.RequestId
	}

	env, err := buildAndSign(r, &xdrbuild.CreateAsset{
		RequestID:                requestID,
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
