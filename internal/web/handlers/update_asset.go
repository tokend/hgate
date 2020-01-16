package handlers

import (
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func UpdateAsset(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewUpdateAssetRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	requestID := uint64(0)
	if request.RequestId != nil {
		requestID = *request.RequestId
	}

	env, err := buildAndSign(r, &xdrbuild.UpdateAsset{
		RequestID:      requestID,
		AllTasks:       request.AllTasks,
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
