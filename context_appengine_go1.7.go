// +build appengine,go1.7

package context

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
)

func getContext(r *http.Request) context.Context {
	return context.WithValue(appengine.NewContext(r), ContextKeyEnvironment, "appengine")
}

// setNamespace sets a custom namespace if the `Namespace` variable is not nil
func setNamespace(ctx context.Context) (context.Context, error) {
	if Namespace == nil {
		return ctx, nil
	}
	ns, err := Namespace(ctx)
	if err != nil {
		return ctx, err
	}
	return appengine.Namespace(ctx, ns)
}

// Hostname returns the hostname of the current instance
func Hostname(ctx context.Context, r *http.Request) (string, error) {
	return appengine.ModuleHostname(ctx, "", "", "")
}

func WithValue(ctx Context, key, val interface{}) Context {
	return context.WithValue(ctx, key, val)
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	return ctx, CancelFunc(cancel)
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	ctx, cancel := context.WithDeadline(parent, deadline)
	return ctx, CancelFunc(cancel)
}
