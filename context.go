package context

import (
	"time"

	"golang.org/x/net/context"
)

// ContextKey is a string type for context value key names
type ContextKey string
type CancelFunc func()

// Context key location definitions
var (
	ContextKeyNamepace       ContextKey = "namespace"
	ContextKeyRequestBody    ContextKey = "request.body"
	ContextKeyRequestVars    ContextKey = "request.vars"
	ContextKeyResponseWriter ContextKey = "http.responsewriter"
	ContextKeyRequest        ContextKey = "request"
	ContextKeyResponseCode   ContextKey = "response.code"
	ContextKeyInitialized    ContextKey = "initialized"
	ContextKeyResponseBody   ContextKey = "response.body"
	ContextKeyEnvironment    ContextKey = "environment"
)

var Namespace = func(ctx Context) (string, error) {
	return "", nil
}

// Context has the same interface as "context" and "golang.org/x/net/context" to handle
// imports in all go version
type Context interface{
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

func SetCode(ctx Context, code int) Context {
	return context.WithValue(ctx, ContextKeyResponseCode, code)
}

