package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type ChangeRoleRequest struct {
	resources.ChangeRoleRequest
}

func (c ChangeRoleRequest) Validate() error {
	errs := validation.Errors{
		"/data/destination":     validation.Validate(c.Destination, validation.Required, validation.Length(56, 56)),
		"/data/account_role":    validation.Validate(c.AccountRole, validation.Required, validation.Min(uint64(1))),
		"/data/creator_details": validation.Validate(c.CreatorDetails, validation.Required),
		"/data/request_id":      validation.Validate(c.CreatorDetails, validation.NilOrNotEmpty),
	}

	return errs.Filter()
}

func NewChangeRoleRequest(r *http.Request) (*ChangeRoleRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request ChangeRoleRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.Validate()

}
