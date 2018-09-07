package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/go/xdrbuild"

	"gitlab.com/tokend/hgate/internal/helpers"
)

type saleDetails struct {
	Description string `json:"description"`
	Logo        struct {
		Key      string `json:"key"`
		MimeType string `json:"mime_type"`
		Name     string `json:"name"`
	} `json:"logo"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	YoutubeVideoID   string `json:"youtube_video_id"`
}
type updateSaleDetails struct {
	Details saleDetails `json:"details"`
}

func ProcessUpdateSaleDetails(r *http.Request) (*xdrbuild.UpdateSaleDetails, error) {
	var request updateSaleDetails
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	saleid, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse sale id")
	}

	out := xdrbuild.UpdateSaleDetails{
		SaleID:         saleid,
		NewSaleDetails: xdrbuild.SaleDetails(request.Details),
	}

	return &out, out.Validate()
}

func UpdateSaleDetails(submitTx helpers.TxSubmitter, logger *logan.Entry) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		op, err := ProcessUpdateSaleDetails(r)
		if err != nil {
			logger.WithError(err).Error("Failed to process request.")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		result, err := submitTx(op)
		if result != nil {
			w.WriteHeader(result.StatusCode)
			w.Write(result.RawResponse)
			return
		}
		if err != nil {
			logger.WithError(err).Error("Failed to submit tx.")
		}
		ape.RenderErr(w, problems.InternalError())
	}
}
