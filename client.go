package geonames

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Client is the main entry point for the geonames library
type Client struct {
	c             http.Client
	cacheDuration time.Duration
}

// NewClient creates a new Client with the specified cache duration
func NewClient(cacheDuration *time.Duration) *Client {
	if cacheDuration != nil {
		return &Client{
			cacheDuration: *cacheDuration,
		}
	}

	return &Client{
		cacheDuration: time.Hour * 24 * 7,
	}
}

func generateFilename(url string) string {
	// Create a new hash.
	h := sha256.New()

	// Write the URL to the hash.
	h.Write([]byte(url))

	// Get the hash as a byte slice.
	hash := h.Sum(nil)

	// Convert the hash to a hexadecimal string.
	filename := hex.EncodeToString(hash)

	return filename
}
func (c *Client) httpGetWithCache(url string) ([]byte, error) {
	// Generate a unique filename based on the URL
	filename := generateFilename(url)

	// Construct the path to the temporary folder and the filename
	tmpDir := os.TempDir()
	filePath := filepath.Join(tmpDir, filename)

	// Check if the file exists at that path
	if info, err := os.Stat(filePath); err == nil {
		// If it does, check the timestamp of the file
		if time.Since(info.ModTime()) < c.cacheDuration {
			// If the timestamp is less than one week old, read the file and return its contents
			file, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			return file, nil
		}
	}

	// If the timestamp is more than one week old or the file doesn't exist, download the file
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Save it to the path
	err = os.WriteFile(filePath, body, 0666)
	if err != nil {
		return nil, err
	}

	// Return its contents
	return body, nil
}
