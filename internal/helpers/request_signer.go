package helpers

import (
	"net/http"

	oldkeypair "gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/signcontrol"
	"gitlab.com/tokend/keypair"
)

type RequestSigner func(r *http.Request) error

func GetRequestSigner(signer keypair.Full) RequestSigner {
	return func(r *http.Request) error {
		return signcontrol.SignRequest(r, oldkeypair.MustParse(signer.Seed()))
	}
}
