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
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string `json:"state,omitempty"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body,omitempty"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//!-
