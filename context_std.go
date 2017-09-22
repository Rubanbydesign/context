// +build !appengine,!go1.7

package context

import (
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"
)

func getContext(r *http.Request) context.Context {
	return context.WithValue(context.Background(), ContextKeyEnvironment, "standard")
}

// Hostname returns the hostname of the current instance
func Hostname(ctx Context, r *http.Request) (string, error) {
	return strings.Split(r.Host, ":")[0], nil
}

func WithValue(ctx Context, key, val interface{}) Context {
	return context.WithValue(ctx, key, val)
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	return context.WithDeadline(parent, deadline)
}
