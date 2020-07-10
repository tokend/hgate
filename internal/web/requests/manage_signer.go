package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type ManageSignerRequest struct {
	resources.ManageSignerRequest
}

func (c ManageSignerRequest) Validate() error {
	errs := validation.Errors{
		"/data/action": validation.Validate(c.Action, validation.Min(0), validation.Max(2)),
	}
	switch c.Action {
	case 0, 1:
		{
			errs["/data/signer_data/weight"] = validation.Validate(c.SignerData.Weight, validation.Required, validation.Min(uint32(1)), validation.Max(uint32(1000)))
			errs["/data/signer_data/public_key"] = validation.Validate(c.SignerData.PublicKey, validation.Required, validation.Length(56, 56))
			errs["/data/signer_data/identity"] = validation.Validate(c.SignerData.Identity, validation.Required, validation.Min(uint32(1)))
			errs["/data/signer_data/signer_role"] = validation.Validate(c.SignerData.SignerRole, validation.Required, validation.Min(uint64(1)))
			errs["/data/signer_data/details"] = validation.Validate(c.SignerData.Details, validation.Required)
		}
	case 2:
		{
			errs["/data/remove_signer_data/public_key"] = validation.Validate(c.RemoveSignerData.PublicKey, validation.Required, validation.Length(56, 56))
		}
	}

	return errs.Filter()
}

func NewManageSignerRequest(r *http.Request) (*ManageSignerRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request ManageSignerRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.Validate()

}
