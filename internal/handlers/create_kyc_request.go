package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"

	"gitlab.com/tokend/hgate/internal/helpers"
)

type сreateUpdateKYCRequest struct {
	RequestID          uint64  `json:"request_id"`
	AccountToUpdateKYC string  `json:"account_to_update_kyc"`
	AccountTypeToSet   string  `json:"account_type_to_set"`
	KYCLevelToSet      uint32  `json:"kyc_level_to_set"`
	KYCData            string  `json:"kyc_data"`
	AllTasks           *uint32 `json:"all_tasks"`
}

func ProcessCreateKYCRequest(r *http.Request) (*xdrbuild.CreateUpdateKYCRequestOp, error) {
	var request сreateUpdateKYCRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var finalGuess xdr.AccountType
	for _, guess := range xdr.AccountTypeAll {
		if guess.ShortString() == request.AccountTypeToSet {
			finalGuess = guess
			break
		}
	}

	out := xdrbuild.CreateUpdateKYCRequestOp{
		RequestID:          request.RequestID,
		AccountToUpdateKYC: request.AccountToUpdateKYC,
		AccountTypeToSet:   finalGuess,
		KYCLevelToSet:      request.KYCLevelToSet,
		KYCData:            request.KYCData,
		AllTasks:           request.AllTasks,
	}

	return &out, out.Validate()
}

func CreateKYCRequest(submitTx helpers.TxSubmitter, logger *logan.Entry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		op, err := ProcessCreateKYCRequest(r)
		if err != nil {
			logger.WithError(err).Error("Failed to process request.")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		result, err := submitTx(op)
		if result != nil {
			w.WriteHeader(result.StatusCode)
			w.Write(result.RawResponse)
			return
		}
		if err != nil {
			logger.WithError(err).Error("Failed to submit tx.")
		}
		ape.RenderErr(w, problems.InternalError())
	}
}
