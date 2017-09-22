package context

import (
	"time"

	"golang.org/x/net/context"
)

// Key is a string type for context value key names
type Key string

// CancelFunc is a context cancel func, matching the "context".CancelFunc interface
type CancelFunc func()

// Context key location definitions
var (
	ContextKeyNamepace       Key = "namespace"
	ContextKeyRequestBody    Key = "request.body"
	ContextKeyRequestVars    Key = "request.vars"
	ContextKeyResponseWriter Key = "http.responsewriter"
	ContextKeyRequest        Key = "request"
	ContextKeyResponseCode   Key = "response.code"
	ContextKeyInitialized    Key = "initialized"
	ContextKeyResponseBody   Key = "response.body"
	ContextKeyEnvironment    Key = "environment"
)

// Namespace is a overrideable function which returns the namespace associated with the given context.
var Namespace = func(ctx Context) (string, error) {
	return "", nil
}

// Context has the same interface as "context" and "golang.org/x/net/context" to handle
// imports in all go version
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

// SetCode sets the HTTP response code associated with the context
func SetCode(ctx Context, code int) Context {
	return context.WithValue(ctx, ContextKeyResponseCode, code)
}
