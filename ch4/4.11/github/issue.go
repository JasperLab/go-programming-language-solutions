package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var githubToken string

func init() {
	githubToken = os.Getenv("GITHUB_TOKEN")
	if len(githubToken) == 0 {
		log.Fatal("$GITHUB_TOKEN not set")
	}
}

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

func CreateIssue(owner string, repo string, issue *Issue) (*Issue, error) {
	url := fmt.Sprintf(IssuesURL, owner, repo)
	return upsertIssue(issue, owner, repo, url, http.MethodPost)
}

func UpdateIssue(issue *Issue, owner string, repo string) (*Issue, error) {
	if issue.Number < 1 {
		return nil, errors.New("Invalid issue number")
	}
	issue_id := strconv.Itoa(issue.Number)
	url := fmt.Sprintf(IssueURL, owner, repo, issue_id)
	return upsertIssue(issue, owner, repo, url, http.MethodPatch)
}

func upsertIssue(issue *Issue, owner string, repo string, url string, httpMethod string) (*Issue, error) {
	payload, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+githubToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("Create/Update issue call failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}
