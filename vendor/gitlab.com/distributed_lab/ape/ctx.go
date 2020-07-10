package ape

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
)

func setContextLog(ctx context.Context, log *logan.Entry) context.Context {
	return context.WithValue(ctx, logCtxKey, log)
}

func getContextLog(ctx context.Context) *logan.Entry {
	return ctx.Value(logCtxKey).(*logan.Entry)
}

// Log allows to retrieve request fielded logan.Entry,
// useful only with DefaultMiddlewares.
func Log(r *http.Request) *logan.Entry {
	return getContextLog(r.Context())
}
