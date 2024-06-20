package auth

import (
	"net/http"

	"github.com/Etwodev/Phasor/crypto"
)

func PubKeyGetRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)

	publicKey, err := crypto.GetPublicKey()
	if err != nil {
		http.Error(w, "Missing public key", http.StatusInternalServerError)
	}

	w.Write(publicKey)
}

func PingGetRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func RedirectGetRoute(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusFound)
}
