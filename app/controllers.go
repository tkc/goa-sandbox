// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa sample": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/tkc/goa-sandbox/design
// --out=$(GOPATH)/src/github.com/tkc/goa-sandbox
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AccountController is the controller interface for the Account actions.
type AccountController interface {
	goa.Muxer
	CurrentUser(*CurrentUserAccountContext) error
	List(*ListAccountContext) error
	Login(*LoginAccountContext) error
}

// MountAccountController "mounts" a Account resource controller on the given service.
func MountAccountController(service *goa.Service, ctrl AccountController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/current_user", ctrl.MuxHandler("preflight", handleAccountOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/list", ctrl.MuxHandler("preflight", handleAccountOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/login", ctrl.MuxHandler("preflight", handleAccountOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCurrentUserAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.CurrentUser(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("GET", "/current_user", ctrl.MuxHandler("currentUser", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "CurrentUser", "route", "GET /current_user")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AccountListPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.List(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("GET", "/list", ctrl.MuxHandler("list", h, unmarshalListAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "List", "route", "GET /list")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginAccountContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Login(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("POST", "/login", ctrl.MuxHandler("login", h, unmarshalLoginAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "Login", "route", "POST /login")
}

// handleAccountOrigin applies the CORS response headers corresponding to the origin.
func handleAccountOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalListAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalListAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &accountListPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalLoginAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalLoginAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &accountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// JWTController is the controller interface for the JWT actions.
type JWTController interface {
	goa.Muxer
	Secure(*SecureJWTContext) error
	SignIn(*SignInJWTContext) error
	Unsecure(*UnsecureJWTContext) error
}

// MountJWTController "mounts" a JWT resource controller on the given service.
func MountJWTController(service *goa.Service, ctrl JWTController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/jwt", ctrl.MuxHandler("preflight", handleJWTOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/jwt/sign_in", ctrl.MuxHandler("preflight", handleJWTOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/jwt/unsecure", ctrl.MuxHandler("preflight", handleJWTOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSecureJWTContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Secure(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleJWTOrigin(h)
	service.Mux.Handle("GET", "/jwt", ctrl.MuxHandler("secure", h, nil))
	service.LogInfo("mount", "ctrl", "JWT", "action", "Secure", "route", "GET /jwt", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSignInJWTContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.SignIn(rctx)
	}
	h = handleSecurity("SigninBasicAuth", h)
	h = handleJWTOrigin(h)
	service.Mux.Handle("POST", "/jwt/sign_in", ctrl.MuxHandler("signIn", h, nil))
	service.LogInfo("mount", "ctrl", "JWT", "action", "SignIn", "route", "POST /jwt/sign_in", "security", "SigninBasicAuth")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUnsecureJWTContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Unsecure(rctx)
	}
	h = handleJWTOrigin(h)
	service.Mux.Handle("GET", "/jwt/unsecure", ctrl.MuxHandler("unsecure", h, nil))
	service.LogInfo("mount", "ctrl", "JWT", "action", "Unsecure", "route", "GET /jwt/unsecure")
}

// handleJWTOrigin applies the CORS response headers corresponding to the origin.
func handleJWTOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080/swagger") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
