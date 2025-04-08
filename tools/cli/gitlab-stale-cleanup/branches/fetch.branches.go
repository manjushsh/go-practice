package branches

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Branch struct represents a GitLab branch
type Branch struct {
	Name      string `json:"name"`
	Commit    Commit `json:"commit"`
	Protected bool   `json:"protected"`
}

// Commit struct represents commit details in a branch
type Commit struct {
	Date string `json:"created_at"`
}

// FetchBranches fetches all branches for a GitLab project
func FetchBranches(projectID, privateToken string) ([]Branch, error) {
	url := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/repository/branches", projectID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("PRIVATE-TOKEN", privateToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var branches []Branch
	if err := json.NewDecoder(resp.Body).Decode(&branches); err != nil {
		return nil, err
	}

	return branches, nil
}
