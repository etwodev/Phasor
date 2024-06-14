package auth

import (
	"net/http"

	"github.com/Etwodev/Phasor/crypto"
)

func TestGetRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)

	publicKey, err := crypto.GetPublicKey()
	if err != nil {
		http.Error(w, "Missing public key", http.StatusInternalServerError)
	}

	w.Write(publicKey)
}
