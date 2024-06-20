package redirect

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Etwodev/ramchi/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

type Redirects struct {
	Redirects map[string]Redirect `bson:"redirects"`
}

type Redirect struct {
	Out       string `bson:"out"`
	CreatedAt string `bson:"created_at"`
}

func RedirectGetRoute(w http.ResponseWriter, r *http.Request) {
	value := helpers.URLParam(r, "id")
	if value == "" {
		http.Error(w, "Bad ID", http.StatusBadRequest)
		return
	}

	// Load existing redirects
	var existingRedirects Redirects
	err := DecodeBSONFromFile("./.redirects", &existingRedirects)
	if err != nil {
		http.Error(w, "Failed to load redirects", http.StatusInternalServerError)
		return
	}

	// Check if the ID exists
	redirect, exists := existingRedirects.Redirects[value]
	if !exists {
		http.Error(w, "ID doesn't exist", http.StatusNotFound)
		return
	}

	// Redirect to the URL if the ID exists
	http.Redirect(w, r, redirect.Out, http.StatusFound)
}

func RedirectPostRoute(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Convert body to string
	url := string(body)

	token := GenerateToken(6)

	redirectItem := Redirect{
		Out:       url,
		CreatedAt: time.Now().Format(time.RFC3339), // Using RFC3339 format for the created_at field
	}

	// Load existing redirects
	var existingRedirects Redirects
	err = DecodeBSONFromFile("./.redirects", &existingRedirects)
	if err != nil {
		// If there's an error because the file doesn't exist, initialize the map
		existingRedirects.Redirects = make(map[string]Redirect)
	}

	// Add the new redirect to the existing ones
	existingRedirects.Redirects[token] = redirectItem

	// Save the updated redirects back to the file
	err = EncodeBSONToFile(&existingRedirects, "./.redirects")
	if err != nil {
		http.Error(w, "Failed to save redirect", http.StatusInternalServerError)
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"data": fmt.Sprintf("%v", existingRedirects), "url": "https://link.etwo.dev/" + token})
}

func GenerateToken(length int) string {
	// Generate random bytes
	buff := make([]byte, length)
	if _, err := rand.Read(buff); err != nil {
		return ""
	}

	// Hash the combined bytes to ensure a fixed length and mix in the timestamp
	encodedString := base64.URLEncoding.EncodeToString(buff[:])

	return encodedString
}

// EncodeBSONToFile encodes a BSON document and writes it to a file
func EncodeBSONToFile(data *Redirects, filePath string) error {
	// Marshal the data to BSON
	bsonData, err := bson.Marshal(&data)
	if err != nil {
		return fmt.Errorf("failed to marshal data to BSON: %v", err)
	}

	// Create and open the file for writing
	err = os.WriteFile(filePath, bsonData, 0644) // 0644 is commonly used permissions for read & write
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

// DecodeBSONFromFile reads a BSON file and decodes it into the provided data structure
func DecodeBSONFromFile(filePath string, result *Redirects) error {
	// Open the file for reading
	bsonData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Unmarshal the BSON data into the provided result structure
	err = bson.Unmarshal(bsonData, &result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal BSON data: %v", err)
	}

	return nil
}
