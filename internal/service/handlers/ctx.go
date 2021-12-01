package handlers

import (
    "context"
    "gitlab.com/tokend/subgroup/tokenproject/internal/data"
    "net/http"

    "gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
    logCtxKey ctxKey = iota
    tokenCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
    return func(ctx context.Context) context.Context {
        return context.WithValue(ctx, logCtxKey, entry)
    }
}

func Log(r *http.Request) *logan.Entry {
    return r.Context().Value(logCtxKey).(*logan.Entry)
}

func Token(r *http.Request) data.TokenQ {
    return r.Context().Value(tokenCtxKey).(data.TokenQ).New()
}

func CtxToken(q data.TokenQ) func(context.Context) context.Context {
    return func(ctx context.Context) context.Context {
        return context.WithValue(ctx, tokenCtxKey, q)
    }
}
