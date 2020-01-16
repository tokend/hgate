package submit

import (
	"context"
	"encoding/json"
	"gitlab.com/distributed_lab/json-api-connector/cerrors"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/regources/generated"
	"net/http"
)

var (
	//ErrSubmitTimeout indicates that transaction submission has timed out
	ErrSubmitTimeout = errors.New("submit timed out")
	//ErrSubmitInternal indicates that transaction submission has failed with internal error
	ErrSubmitInternal = errors.New("internal submit error")
	//ErrSubmitUnexpectedStatusCode indicates that transaction submission has failed with unexpected status code
	ErrSubmitUnexpectedStatusCode = errors.New("unexpected unsuccessful status code")
)

func (t *Transactor) Submit(ctx context.Context, envelope string, waitIngest bool) (*regources.TransactionResponse, error) {
	body := regources.SubmitTransactionBody{
		Tx:            envelope,
		WaitForIngest: &waitIngest,
	}

	var success regources.TransactionResponse

	err := t.base.PostJSON(t.submissionUrl, body, ctx, &success)
	if err == nil {
		return &success, nil
	}

	cerr, ok := err.(cerrors.Error)
	if !ok {
		return nil, errors.Wrap(err, "failed to submit tx to horizon")
	}

	// go through known response codes and try to build meaningful result
	switch cerr.Status() {
	case http.StatusGatewayTimeout: // timeout
		return nil, ErrSubmitTimeout
	case http.StatusBadRequest: // rejected or malformed
		// check which error it was exactly, might be useful for consumer
		var failureResp txFailureResponse
		if err := json.Unmarshal(cerr.Body(), &failureResp); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal horizon response")
		}
		return nil, newTxFailure(failureResp)
	case http.StatusInternalServerError: // internal error
		return nil, ErrSubmitInternal
	default:
		return nil, errors.Wrap(err, ErrSubmitUnexpectedStatusCode.Error())
	}
}
