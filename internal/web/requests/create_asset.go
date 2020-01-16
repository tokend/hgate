package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type CreateAssetRequest struct {
	resources.CreateAssetRequest
}

func (c CreateAssetRequest) Validate() error {
	errs := validation.Errors{
		"/data/request_id": validation.Validate(c.RequestId, validation.NilOrNotEmpty),
		//todo add regular expression matching for asset code
		"/data/code":                     validation.Validate(c.Code, validation.Length(1, 16)),
		"/data/type":                     validation.Validate(c.Type, validation.Required),
		"/data/creator_details":          validation.Validate(c.CreatorDetails, validation.Required),
		"/data/initial_preissued_amount": validation.Validate(c.InitialPreissuedAmount, validation.Required),
		"/data/max_issuance_amount":      validation.Validate(c.MaxIssuanceAmount, validation.Required),
		"/data/trailing_digits_count":    validation.Validate(c.TrailingDigitsCount, validation.Required, validation.Min(uint32(0)), validation.Max(uint32(6))),
		"/data/preissuance_asset_signer": validation.Validate(c.PreIssuanceAssetSigner, validation.Required, validation.Length(56, 56)),
		"/data/policies":                 validation.Validate(c.Policies, validation.Required, validation.Max(uint32(255))),
	}

	return errs.Filter()
}

func NewCreateAssetRequest(r *http.Request) (*CreateAssetRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request CreateAssetRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.Validate()

}
