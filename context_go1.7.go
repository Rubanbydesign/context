// +build !appengine,go1.7

package context

import (
	"net/http"
	"strings"
	"time"

	"context"
)

func getContext(r *http.Request) context.Context {
	return context.WithValue(context.Background(), ContextKeyEnvironment, "standard")
}

// setNamespace sets a custom namespace if the `Namespace` variable is not nil
func setNamespace(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// NewContext creates a new context associated with the request.
func NewContext(r *http.Request) Context {
	return context.Background()
}

// Hostname returns the hostname of the current instance
func Hostname(ctx Context, r *http.Request) (string, error) {
	return strings.Split(r.Host, ":")[0], nil
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(ctx Context, key, val interface{}) Context {
	return context.WithValue(ctx, key, val)
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	return ctx, CancelFunc(cancel)
}

// WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d.
// If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent
// to parent. The returned context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is closed, whichever happens first.
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	ctx, cancel := context.WithDeadline(parent, deadline)
	return ctx, CancelFunc(cancel)
}

// Background returns a new background context
func Background() Context {
	return context.Background()
}
