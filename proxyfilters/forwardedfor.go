package proxyfilters

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/getlantern/proxy/filters"
)

const (
	xForwardedFor = "X-Forwarded-For"
)

// AddForwardedFor adds an X-Forwarded-For header based on the request's
// RemoteAddr.
var AddForwardedFor = filters.FilterFunc(func(ctx context.Context, req *http.Request, next filters.Next) (*http.Response, error) {
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := req.Header[xForwardedFor]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		req.Header.Set(xForwardedFor, clientIP)
	}
	return next(ctx, req)
})