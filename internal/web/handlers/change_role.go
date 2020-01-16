package handlers

import (
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func ChangeRole(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewChangeRoleRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	requestID := uint64(0)
	if request.RequestId != nil {
		requestID = *request.RequestId
	}

	kycData := make(map[string]interface{})
	err = json.Unmarshal(request.CreatorDetails, &kycData)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"/data/creator_details": errors.New("invalid creator details: creator details must be a valid JSON object"),
		})...)
	}

	env, err := buildAndSign(r, &xdrbuild.CreateChangeRoleRequest{
		RequestID:          requestID,
		DestinationAccount: request.Destination,
		RoleToSet:          request.AccountRole,
		KYCData:            kycData,
		AllTasks:           request.AllTasks,
	})

	if err != nil {
		log.WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}
