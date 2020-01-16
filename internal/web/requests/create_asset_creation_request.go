package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type CreateAssetCreationRequest AssetCreationRequest

func NewCreateAssetCreationRequest(r *http.Request) (*CreateAssetCreationRequest, error) {
	var d dummy
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	var request CreateAssetCreationRequest
	if err := json.Unmarshal(d.Data, &request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, AssetCreationRequest(request).Validate()
}
