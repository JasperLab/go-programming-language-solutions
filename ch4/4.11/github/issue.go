package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func UpdateIssue(issue *Issue, owner string, token string, repo string) (*Issue, error) {
	payload, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}

	issue_id := strconv.Itoa(issue.Number)
	url := fmt.Sprintf(IssueURL, owner, repo, issue_id)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Update issue call failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func CreateIssue(owner string, repo string, description string) (*Issue, error) {
	return nil, nil
}
