package handlers

import (
	"context"
	"github.com/tokend/hgate/internal/helpers"
	"gitlab.com/tokend/connectors/submit"
	oldkeypair "gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/signcontrol"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	builderCtxKey
	keysCtxKey
	submitterCtxKey
	proxyCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxBuilder(builder *xdrbuild.Builder) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, builderCtxKey, builder)
	}
}

func Builder(r *http.Request) *xdrbuild.Builder {
	return r.Context().Value(builderCtxKey).(*xdrbuild.Builder)
}

func CtxKeys(keys SigningBundle) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, keysCtxKey, keys)
	}
}

func Keys(r *http.Request) SigningBundle {
	return r.Context().Value(keysCtxKey).(SigningBundle)
}

func CtxSubmitter(transactor *submit.Transactor) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, submitterCtxKey, transactor)
	}
}

func Submitter(r *http.Request) *submit.Transactor {
	return r.Context().Value(submitterCtxKey).(*submit.Transactor)
}

func signRequest(r *http.Request, signer keypair.Full) error {
	return signcontrol.SignRequest(r, oldkeypair.MustParse(signer.Seed()))
}

func CtxProxy(proxy helpers.Proxy) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, proxyCtxKey, proxy)
	}
}

func Proxy(w http.ResponseWriter, r *http.Request) {
	proxy := r.Context().Value(proxyCtxKey).(helpers.Proxy)
	proxy(w, r)
}
