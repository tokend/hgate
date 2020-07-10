package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/tokend/hgate/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type PaymentRequest struct {
	resources.PaymentRequest
}

func (c PaymentRequest) Validate() error {
	errs := validation.Errors{
		"/data/destination_balance": validation.Validate(c.DestinationBalance, validation.NilOrNotEmpty, validation.Length(56, 56)),
		"/data/reference":           validation.Validate(c.Reference, validation.Required),
		"/data/amount":              validation.Validate(c.Amount, validation.Required),
		"/data/source_balance":      validation.Validate(c.SourceBalance, validation.Required, validation.Length(56, 56)),
	}

	if c.DestinationBalance == nil {
		errs["/data/destination_account"] = validation.Validate(c.DestinationAccount, validation.Required, validation.Length(56, 56))

	}

	return errs.Filter()
}

func NewPaymentRequest(r *http.Request) (*PaymentRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request PaymentRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.Validate()

}
