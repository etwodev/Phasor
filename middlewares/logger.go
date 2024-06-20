package middlewares

import (
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02T15:04:05"}).
	With().
	Timestamp().
	Str("Group", "phasor").
	Logger()

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			duration := time.Since(start)

			log.Info().
				Str("IP", getIPAddress(r)).
				Str("Url", r.URL.Path).
				Str("Method", r.Method).
				Dur("response_time", duration).
				Msg("Incoming request")
		}
		return http.HandlerFunc(fn)
	}
}

func getIPAddress(r *http.Request) string {
	// Try to get the IP from the X-Forwarded-For header
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// This header can contain multiple IPs. The first one is the original client IP
		parts := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(parts[0])
	}

	// If there is no X-Forwarded-For header, use the RemoteAddr from the request
	// This might include a port, so we split it off if present
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// If there was an error, it might be because there's no port in the RemoteAddr
		// In that case, just return the RemoteAddr as the IP
		return r.RemoteAddr
	}
	return ip
}
