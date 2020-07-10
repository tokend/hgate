package submit

import (
	connector "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/json-api-connector/client"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdrbuild"
	regources "gitlab.com/tokend/regources/generated"
	"net/url"
)

type Transactor struct {
	base *connector.Connector

	infoUrl       *url.URL
	submissionUrl *url.URL
}

func New(client client.Client) *Transactor {
	info, _ := url.Parse("/v3/info")
	submission, _ := url.Parse("/v3/transactions")

	return &Transactor{
		base:          connector.NewConnector(client),
		infoUrl:       info,
		submissionUrl: submission,
	}
}

func (t *Transactor) TXBuilder() (*xdrbuild.Builder, error) {
	info, err := t.info()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get horizon info")
	}

	return xdrbuild.NewBuilder(info.Attributes.NetworkPassphrase, info.Attributes.TxExpirationPeriod), nil
}

func (t *Transactor) info() (*regources.HorizonState, error) {
	var resp regources.HorizonStateResponse

	err := t.base.Get(t.infoUrl, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	return &resp.Data, nil
}
