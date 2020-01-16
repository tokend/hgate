package handlers

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/connectors/keyer"
	"gitlab.com/tokend/connectors/submit"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
	"net/http"
)

func proxyTransaction(r *http.Request, w http.ResponseWriter, envelope string) {
	resp, err := Submitter(r).Submit(r.Context(), envelope, true)
	if err != nil {

		if serr, ok := err.(submit.TxFailure); ok {
			w.WriteHeader(http.StatusBadRequest)
			ape.Render(w, serr.Response)
			return
		}

		Log(r).WithError(err).Error("failed to submit transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resp)
}

func buildAndSign(r *http.Request, ops ...xdrbuild.Operation) (string, error) {
	keys, err := getKeys(r)
	if err != nil {
		return "", errors.Wrap(err, "failed to get keys")
	}

	tx := Builder(r).Transaction(keys.Source)
	for _, op := range ops {
		tx = tx.Op(op)
	}

	env, err := tx.Sign(keys.Signer).Marshal()
	return env, err
}

func getKeys(r *http.Request) (keyer.Keys, error) {
	source := r.Header["Tokend-Source"]
	if source == nil || len(source) != 1 {
		return Keys(r), nil
	}
	signer := r.Header["Tokend-Signer"]
	if signer == nil || len(signer) != 1 {
		return Keys(r), nil
	}

	sourceAddr, err := keypair.ParseAddress(source[0])
	if err != nil {
		return keyer.Keys{}, errors.Wrap(err, "failed to parse source")
	}

	signerKey, err := keypair.ParseSeed(signer[0])
	if err != nil {
		return keyer.Keys{}, errors.Wrap(err, "failed to parse signer")
	}

	return keyer.Keys{
		Source: sourceAddr,
		Signer: signerKey,
	}, nil

}
