package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"ripe-data-exporter/pkg/types"
)

func FetchRIPEData(apiURL string, timeout time.Duration) (*types.Response, error) {
	client := &http.Client{Timeout: timeout}

	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status code %d and message: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response types.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &response, nil
}
