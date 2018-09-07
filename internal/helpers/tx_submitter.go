package helpers

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdrbuild"
	horizon "gitlab.com/tokend/horizon-connector"
	"gitlab.com/tokend/keypair"
)

type TxSubmitter func(operation xdrbuild.Operation) (*horizon.SubmitResponseDetails, error)

func GetTxSubmitter(kp keypair.Full, horizonURL *url.URL) (TxSubmitter, error) {
	connector := horizon.NewConnector(horizonURL).WithSigner(kp)
	builder, err := connector.TXBuilder()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create builder")
	}

	return func(operation xdrbuild.Operation) (*horizon.SubmitResponseDetails, error) {
		env, err := builder.Transaction(kp).Op(operation).Sign(kp).Marshal()
		fmt.Println(kp.Address())
		fmt.Println(kp.Seed())
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal tx")
		}

		resp, err := connector.Submitter().SubmitE(env)
		if err != nil {
			return &resp, errors.Wrap(err, "failed to submit tx")
		}

		return &resp, nil
	}, nil
}
