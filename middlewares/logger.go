package middlewares

import (
	"net/http"
	"os"
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
				Str("IP", r.RemoteAddr).
				Str("Url", r.URL.Path).
				Str("Method", r.Method).
				Dur("response_time", duration).
				Msg("Incoming request")
		}
		return http.HandlerFunc(fn)
	}
}
