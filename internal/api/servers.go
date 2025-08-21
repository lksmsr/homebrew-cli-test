package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type CreateServerRequest struct {
	Name         string `json:"name"`
	InstanceType string `json:"instanceType"`
	Location     string `json:"location"`
}

type CreateServerResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InstanceType string `json:"instanceType"`
	Location     string `json:"location"`
	Status       string `json:"status"`
	IPv4         string `json:"ipv4"`
	IPv6         string `json:"ipv6"`
	CreatedAt    string `json:"createdAt"`
}

func (sliplaneApiClient *SliplaneApiClient) CreateServer(name, instanceType, location string) (*CreateServerResponse, error) {
	req := CreateServerRequest{
		Name:         name,
		InstanceType: instanceType,
		Location:     location,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := sliplaneApiClient.Post("servers", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}
	var createResp CreateServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		return nil, err
	}
	return &createResp, nil
}

type ListServersResponse []struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InstanceType string `json:"instanceType"`
	Location     string `json:"location"`
	Status       string `json:"status"`
	IPv4         string `json:"ipv4"`
	IPv6         string `json:"ipv6"`
	CreatedAt    string `json:"createdAt"`
}

func (sliplaneApiClient *SliplaneApiClient) ListServers() (ListServersResponse, error) {
	resp, err := sliplaneApiClient.Get("servers")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}
	var listResp ListServersResponse

	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp, nil
}

// DeleteServer deletes a server by its ID
func (sliplaneApiClient *SliplaneApiClient) DeleteServer(serverID string) error {
	resp, err := sliplaneApiClient.Delete("servers/" + serverID)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(b))
	}
	return nil
}
