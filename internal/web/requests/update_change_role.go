package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type UpdateChangeRoleRequest struct {
	ChangeRoleRequest
	RequestID uint64
}

func NewUpdateChangeRoleRequest(r *http.Request) (*UpdateChangeRoleRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request UpdateChangeRoleRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	b, err := newBase(r, baseOpts{})
	if err != nil {
		return nil, err
	}
	request.RequestID, err = b.getUint64("request_id")
	if err != nil {
		return nil, err
	}

	if request.AllTasks != nil {
		return nil, validation.Errors{
			"/data/all_tasks": errors.New("Tasks are not allowed on request update"),
		}
	}

	return &request, request.Validate()

}
