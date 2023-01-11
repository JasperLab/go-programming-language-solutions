// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

const SearchURL = "https://api.github.com/search/issues"
const IssueURL = "https://api.github.com/repos/%s/%s/issues/%s"
const IssuesURL = "https://api.github.com/repos/%s/%s/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int    `json:"id,omitempty"`
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string `json:"state,omitempty"`
	User      *User
	CreatedAt time.Time   `json:"created_at"`
	Body      string      `json:"body,omitempty"`
	Assignees []*Assignee `json:"assignees,omitempty"`
	Labels    []*Label    `json:"labels,omitempty"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url,omitempty"`
}

type Label struct {
	Id          int    `json:"id,omitempty"`
	Url         string `json:"url,omitempty"`
	Name        string `json:"name"`
	Description string `json:description,omitempty"`
}

type Assignee struct {
	Login string `json:"login"`
	Id    int    `json:id,omitempty"`
}

//!-
