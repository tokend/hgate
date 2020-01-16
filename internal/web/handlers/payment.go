package handlers

import (
	"github.com/pkg/errors"
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func Payment(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	request, err := requests.NewPaymentRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	sourceBalanceId := xdr.BalanceId{}
	if err := sourceBalanceId.SetString(request.SourceBalance); err != nil {
		log.WithError(err).Error("failed to parse source balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	dest, err := getPaymentDestination(*request)
	if err != nil {
		log.WithError(err).Error("failed to set destination")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	env, err := buildAndSign(r, &xdrbuild.Payment{
		SourceBalanceID: sourceBalanceId,
		Destination:     dest,
		FeeData: xdr.PaymentFeeData{
			SourceFee: xdr.Fee{
				Fixed:   xdr.Uint64(request.SourceFee.Fixed),
				Percent: xdr.Uint64(request.SourceFee.CalculatedPercent),
			},
			DestinationFee: xdr.Fee{
				Fixed:   xdr.Uint64(request.DestinationFee.Fixed),
				Percent: xdr.Uint64(request.DestinationFee.CalculatedPercent),
			},
		},
		Amount:    uint64(request.Amount),
		Subject:   request.Subject,
		Reference: request.Reference,
	})

	if err != nil {
		log.WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}

func getPaymentDestination(request requests.PaymentRequest) (xdr.PaymentOpDestination, error) {
	if request.DestinationBalance != nil {
		balanceId := xdr.BalanceId{}
		if err := balanceId.SetString(*request.DestinationBalance); err != nil {
			return xdr.PaymentOpDestination{}, errors.Wrap(err, "failed to set balance id")
		}

		return xdr.PaymentOpDestination{
			Type:      xdr.PaymentDestinationTypeBalance,
			BalanceId: &balanceId,
		}, nil
	}

	if request.DestinationAccount == nil {
		return xdr.PaymentOpDestination{}, errors.New("destination is missing")
	}

	accID := xdr.AccountId{}
	if err := accID.SetAddress(*request.DestinationAccount); err != nil {
		return xdr.PaymentOpDestination{}, errors.Wrap(err, "failed to set account address")
	}

	return xdr.PaymentOpDestination{
		Type:      xdr.PaymentDestinationTypeAccount,
		AccountId: &accID,
	}, nil

}
