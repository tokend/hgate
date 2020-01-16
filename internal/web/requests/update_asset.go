package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
)

type AssetUpdateRequest struct {
	resources.UpdateAssetRequest
}

func (c AssetUpdateRequest) Validate() error {
	errs := validation.Errors{
		"/data/request_id": validation.Validate(c.RequestId, validation.NilOrNotEmpty),
		//todo add regular expression matching for asset code
		"/data/asset":           validation.Validate(c.Code, validation.Length(1, 16)),
		"/data/creator_details": validation.Validate(c.CreatorDetails, validation.Required),
		"/data/policies":        validation.Validate(c.Policies, validation.Required, validation.Max(uint32(255))),
	}

	return errs.Filter()
}
