package handlers

import (
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func UpdateAssetUpdateRequest(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewUpdateAssetUpdateRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	env, err := buildAndSign(r, &xdrbuild.UpdateAsset{
		RequestID:      request.RequestID,
		Code:           request.Code,
		Policies:       request.Policies,
		CreatorDetails: request.CreatorDetails,
	})

	if err != nil {
		log.WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}
