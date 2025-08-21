package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type CreateProjectResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (sliplaneApiClient *SliplaneApiClient) CreateProject(name string) (*CreateProjectResponse, error) {
	reqBody := CreateProjectRequest{Name: name}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := sliplaneApiClient.Post("projects", bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}
	var projectResp CreateProjectResponse
	if err := json.NewDecoder(resp.Body).Decode(&projectResp); err != nil {
		return nil, err
	}
	return &projectResp, nil
}

type ListProjectsResponse []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (sliplaneApiClient *SliplaneApiClient) ListProjects() (ListProjectsResponse, error) {
	resp, err := sliplaneApiClient.Get("projects")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}
	var listResp ListProjectsResponse

	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp, nil
}

func (sliplaneApiClient *SliplaneApiClient) DeleteProject(projectID string) error {
	resp, err := sliplaneApiClient.Delete("projects/" + projectID)
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
