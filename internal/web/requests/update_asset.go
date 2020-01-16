package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type UpdateAssetRequest struct {
	resources.UpdateAssetRequest
}

func (c UpdateAssetRequest) Validate() error {
	errs := validation.Errors{
		"/data/request_id": validation.Validate(c.RequestId, validation.NilOrNotEmpty),
		//todo add regular expression matching for asset code
		"/data/asset":                    validation.Validate(c.Code, validation.Length(1, 16)),
		"/data/creator_details":          validation.Validate(c.CreatorDetails, validation.Required),
		"/data/policies":                 validation.Validate(c.Policies, validation.Required, validation.Max(uint32(255))),
	}

	return errs.Filter()
}

func NewUpdateAssetRequest(r *http.Request) (*UpdateAssetRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request UpdateAssetRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.Validate()

}
