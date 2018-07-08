package main

import (
	"context"
	"net/http"
	"github.com/goadesign/goa"
)

func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			user, pass, ok := req.BasicAuth()
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				goa.NewErrorClass("unauthorized", 401)
			}
			goa.LogInfo(ctx, "basic", "user", user, "pass", pass)
			return h(ctx, rw, req)
		}
	}
}