// +build go1.7

package context

import (
	"net/http"
	"strings"

	"context"
)

func getContext(r *http.Request) context.Context {
	return context.WithValue(context.Background(), ContextKeyEnvironment, "standard")
}

// setNamespace sets a custom namespace if the `Namespace` variable is not nil
func setNamespace(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func NewContext(r *http.Request) Context {
	return context.Background()
}

// Hostname returns the hostname of the current instance
func Hostname(ctx Context) (string, error) {
	return strings.Split(Request(ctx).Host, ":")[0], nil
}

func WithValue(ctx Context) Context {
	return context.WithValue(ctx, key, val)
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	return context.WithDeadline(parent, deadline)
}