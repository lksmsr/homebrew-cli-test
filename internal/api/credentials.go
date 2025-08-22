package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type ListCredentialsResponse []struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

func (sliplaneApiClient *SliplaneApiClient) ListCredentials() (ListCredentialsResponse, error) {
	resp, err := sliplaneApiClient.Get("registry-credentials")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}
	var listResp ListCredentialsResponse

	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp, nil
}
