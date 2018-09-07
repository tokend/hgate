package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/go/xdrbuild"

	"gitlab.com/tokend/hgate/internal/helpers"
)

type assetDetails struct {
	ExternalSystemType int    `json:"external_system_type,string"`
	Name               string `json:"name"`
}

type updateAsset struct {
	Code     string       `json:"-"`
	Policies uint32       `json:"policies"`
	Details  assetDetails `json:"details"`
}

func ProcessUpdateAssetRequest(r *http.Request) (*xdrbuild.UpdateAsset, error) {
	var request updateAsset
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	request.Code = chi.URLParam(r, "code")

	out := xdrbuild.UpdateAsset{
		Code:     request.Code,
		Policies: request.Policies,
		Details: xdrbuild.AssetDetails{
			ExternalSystemType: request.Details.ExternalSystemType,
			Name:               request.Details.Name,
		},
	}
	return &out, out.Validate()
}

func UpdateAsset(submitTx helpers.TxSubmitter, logger *logan.Entry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		op, err := ProcessUpdateAssetRequest(r)
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
		return
	}
}
