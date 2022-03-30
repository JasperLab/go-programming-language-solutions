package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIssue(owner string, repo string, issue_id string) (*Issue, error) {
	url := fmt.Sprintf(IssueURL, owner, repo, issue_id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Issue query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func UpdateIssue(repo_url string, issue *Issue) (*Issue, error) {
	return nil, nil
}

func CloseIssue(repo_rul string, issue *Issue) (*Issue, error) {
	return nil, nil
}

func CreateIssue(repo_url string) (*Issue, error) {
	return nil, nil
}
