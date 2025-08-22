package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type ListServicesResponse []struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ServerID  string `json:"serverId"`
	ProjectID string `json:"projectId"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	Network   struct {
		Public         bool   `json:"public"`
		Protocol       string `json:"protocol"`
		ManagedDomain  string `json:"managedDomain"`
		InternalDomain string `json:"internalDomain"`
		CustomDomains  []struct {
			ID     string `json:"id"`
			Domain string `json:"domain"`
			Status string `json:"status"`
		} `json:"customDomains"`
	} `json:"network"`
	Volumes []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		MountPath string `json:"mountPath"`
	} `json:"volumes"`
	Env []struct {
		Key    string `json:"key"`
		Value  string `json:"value"`
		Secret bool   `json:"secret"`
	} `json:"env"`
	Deployment struct {
		URL            string `json:"url"`
		DockerfilePath string `json:"dockerfilePath"`
		DockerContext  string `json:"dockerContext"`
		AutoDeploy     bool   `json:"autoDeploy"`
		Branch         string `json:"branch"`
	} `json:"deployment"`
	Healthcheck string `json:"healthcheck"`
	Cmd         string `json:"cmd"`
}

func (sliplaneApiClient *SliplaneApiClient) ListServices(projectID string) (ListServicesResponse, error) {
	resp, err := sliplaneApiClient.Get(fmt.Sprintf("projects/%s/services", projectID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(b))
	}
	var listResp ListServicesResponse

	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, err
	}
	return listResp, nil
}
