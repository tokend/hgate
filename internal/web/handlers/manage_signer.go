package handlers

import (
	"github.com/tokend/hgate/internal/web/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

func ManageSigner(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewManageSignerRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	switch request.Action {
	case 0:
		createSigner(r, w, request)
	case 1:
		updateSigner(r, w, request)
	case 2:
		removeSigner(r, w, request)
	}
}

func updateSigner(r *http.Request, w http.ResponseWriter, request *requests.ManageSignerRequest) {

	signersData := xdrbuild.SignerData{
		PublicKey: request.SignerData.PublicKey,
		RoleID:    request.SignerData.SignerRole,
		Weight:    request.SignerData.Weight,
		Identity:  request.SignerData.Identity,
		Details:   request.SignerData.Details,
	}

	env, err := buildAndSign(r, &xdrbuild.UpdateSigner{
		SignerData: signersData,
	})

	if err != nil {
		Log(r).WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}

func createSigner(r *http.Request, w http.ResponseWriter, request *requests.ManageSignerRequest) {
	signersData := xdrbuild.SignerData{
		PublicKey: request.SignerData.PublicKey,
		RoleID:    request.SignerData.SignerRole,
		Weight:    request.SignerData.Weight,
		Identity:  request.SignerData.Identity,
		Details:   request.SignerData.Details,
	}

	env, err := buildAndSign(r, &xdrbuild.CreateSigner{
		SignerData: signersData,
	})

	if err != nil {
		Log(r).WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}

func removeSigner(r *http.Request, w http.ResponseWriter, request *requests.ManageSignerRequest) {
	env, err := buildAndSign(r, &xdrbuild.RemoveSigner{
		PublicKey: request.RemoveSignerData.PublicKey,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to marshal transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proxyTransaction(r, w, env)
}
