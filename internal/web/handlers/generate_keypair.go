package handlers

import (
	"github.com/tokend/hgate/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/keypair"
	"net/http"
)

func GenerateKeypair(w http.ResponseWriter, r *http.Request) {
	log := Log(r)

	kp, err := keypair.Random()
	if err != nil {
		log.WithError(err).Error("failed to generate keypair")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	resp := resources.KeypairResponse{
		Data: resources.Keypair{
			PublicKey: kp.Address(),
			SecretKey: kp.Seed(),
		},
	}

	ape.Render(w, resp)
}
