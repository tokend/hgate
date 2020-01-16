package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
)

type ChangeRoleRequest resources.ChangeRoleRequest

func (c ChangeRoleRequest) Validate() error {
	errs := validation.Errors{
		"/data/destination":     validation.Validate(c.Destination, validation.Required, validation.Length(56, 56)),
		"/data/account_role":    validation.Validate(c.AccountRole, validation.Required, validation.Min(uint64(1))),
		"/data/creator_details": validation.Validate(c.CreatorDetails, validation.Required),
	}

	return errs.Filter()
}
