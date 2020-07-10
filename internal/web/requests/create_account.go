package requests

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type CreateAccountRequest struct {
	resources.CreateAccountRequest
}

func (c CreateAccountRequest) Validate() error {
	errs := validation.Errors{
		"/data/destination":  validation.Validate(c.Destination, validation.Required, validation.Length(56, 56)),
		"/data/account_role": validation.Validate(c.AccountRole, validation.Required, validation.Min(uint64(1))),
		"/data/referrer":     validation.Validate(c.Referrer, validation.NilOrNotEmpty),
		"/data/signers":      validation.Validate(c.Signers, validation.Required, validation.Length(1, 0)),
	}

	for i, s := range c.Signers {
		errs[fmt.Sprintf("/data/signers/%d/weight", i)] = validation.Validate(s.Weight, validation.Required, validation.Min(uint32(1)), validation.Max(uint32(1000)))
		errs[fmt.Sprintf("/data/signers/%d/public_key", i)] = validation.Validate(s.PublicKey, validation.Required, validation.Length(56, 56))
		errs[fmt.Sprintf("/data/signers/%d/identity", i)] = validation.Validate(s.Identity, validation.Required, validation.Min(uint32(1)))
		errs[fmt.Sprintf("/data/signers/%d/signer_role", i)] = validation.Validate(s.SignerRole, validation.Required, validation.Min(uint64(1)))
		errs[fmt.Sprintf("/data/signers/%d/details", i)] = validation.Validate(s.Details, validation.Required)
	}

	return errs.Filter()
}

func NewCreateAccountRequest(r *http.Request) (*CreateAccountRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request CreateAccountRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.Validate()

}
