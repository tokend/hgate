package handlers

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/connectors/submit"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
	"net/http"
)

type SigningBundle struct {
	Source  keypair.Address
	Signers []keypair.Full
}

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

	for _, s := range keys.Signers {
		tx = tx.Sign(s)
	}

	env, err := tx.Marshal()
	return env, err
}

func getKeys(r *http.Request) (SigningBundle, error) {
	source := r.Header["Tokend-Source"]
	if source == nil || len(source) != 1 {
		return Keys(r), nil
	}
	signer := r.Header["Tokend-Signers"]
	if signer == nil || len(signer) < 1 {
		return Keys(r), nil
	}

	sourceAddr, err := keypair.ParseAddress(source[0])
	if err != nil {
		return SigningBundle{}, errors.Wrap(err, "failed to parse source")
	}

	signingKeys := make([]keypair.Full, 0, len(signer))
	for _, s := range signer {
		signerKey, err := keypair.ParseSeed(s)
		if err != nil {
			return SigningBundle{}, errors.Wrap(err, "failed to parse signer")
		}

		signingKeys = append(signingKeys, signerKey)

	}

	return SigningBundle{
		Source:  sourceAddr,
		Signers: signingKeys,
	}, nil

}
