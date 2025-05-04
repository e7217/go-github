package github

import (
	"context"
	"fmt"
)

type SubIssuesService service

type SubIssue struct {
	ID     *int64  `json:"id,omitempty"`
	Number *int    `json:"number,omitempty"`
	State  *string `json:"state,omitempty"`
	// StateReason can be one of: "completed", "not_planned", "reopened".
	StateReason       *string           `json:"state_reason,omitempty"`
	Locked            *bool             `json:"locked,omitempty"`
	Title             *string           `json:"title,omitempty"`
	Body              *string           `json:"body,omitempty"`
	AuthorAssociation *string           `json:"author_association,omitempty"`
	User              *User             `json:"user,omitempty"`
	Labels            []*Label          `json:"labels,omitempty"`
	Assignee          *User             `json:"assignee,omitempty"`
	Comments          *int              `json:"comments,omitempty"`
	ClosedAt          *Timestamp        `json:"closed_at,omitempty"`
	CreatedAt         *Timestamp        `json:"created_at,omitempty"`
	UpdatedAt         *Timestamp        `json:"updated_at,omitempty"`
	ClosedBy          *User             `json:"closed_by,omitempty"`
	URL               *string           `json:"url,omitempty"`
	HTMLURL           *string           `json:"html_url,omitempty"`
	CommentsURL       *string           `json:"comments_url,omitempty"`
	EventsURL         *string           `json:"events_url,omitempty"`
	LabelsURL         *string           `json:"labels_url,omitempty"`
	RepositoryURL     *string           `json:"repository_url,omitempty"`
	Milestone         *Milestone        `json:"milestone,omitempty"`
	PullRequestLinks  *PullRequestLinks `json:"pull_request,omitempty"`
	Repository        *Repository       `json:"repository,omitempty"`
	Reactions         *Reactions        `json:"reactions,omitempty"`
	Assignees         []*User           `json:"assignees,omitempty"`
	NodeID            *string           `json:"node_id,omitempty"`
	Draft             *bool             `json:"draft,omitempty"`
	Type              *IssueType        `json:"type,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://docs.github.com/rest/search/#text-match-metadata
	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	// ActiveLockReason is populated only when LockReason is provided while locking the issue.
	// Possible values are: "off-topic", "too heated", "resolved", and "spam".
	ActiveLockReason *string `json:"active_lock_reason,omitempty"`
}

// AddSubIssueRequest represents a request to add a sub-issue.
type AddSubIssueRequest struct {
	SubIssueID int `json:"sub_issue_id"`
}

// Add a sub-issue to an issue.
//
// GitHub API docs: https://docs.github.com/en/rest/issues/sub-issues?apiVersion=2022-11-28#add-sub-issue
//
//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/sub_issues
func (s *SubIssuesService) Add(ctx context.Context, owner string, repo string, issueNumber int, subIssueID int) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues", owner, repo, issueNumber)
	req, err := s.client.NewRequest("POST", u, &AddSubIssueRequest{
		SubIssueID: subIssueID,
	})
	if err != nil {
		return nil, err
	}

	// TODO: remove custom Accept header when this API fully launch.
	req.Header.Set("Accept", mediaTypeReactionsPreview)

	return s.client.Do(ctx, req, nil)
}

// Remove a sub-issue from an issue.
//
// GitHub API docs: https://docs.github.com/en/rest/issues/sub-issues?apiVersion=2022-11-28#remove-sub-issue
//
//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/sub_issues/{sub_issue_id}
func (s *SubIssuesService) Remove(ctx context.Context, owner string, repo string, issueNumber int, subIssueID int) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues/%d", owner, repo, issueNumber, subIssueID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	// TODO: remove custom Accept header when this API fully launch.
	req.Header.Set("Accept", mediaTypeReactionsPreview)

	return s.client.Do(ctx, req, nil)
}

// ReprioritizeSubIssueRequest represents a request to reprioritize a sub-issue.
type ReprioritizeSubIssueRequest struct {
	SubIssueID int `json:"sub_issue_id"`
	AfterID    int `json:"after_id,omitempty"`
}

// Reprioritize a sub-issue in an issue.
//
// GitHub API docs: https://docs.github.com/en/rest/issues/sub-issues?apiVersion=2022-11-28#reprioritize-sub-issue
//
//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/sub_issues/priority
func (s *SubIssuesService) Reprioritize(ctx context.Context, owner string, repo string, issueNumber int, subIssueID int, afterID int) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues/priority", owner, repo, issueNumber)
	req, err := s.client.NewRequest("POST", u, &ReprioritizeSubIssueRequest{
		SubIssueID: subIssueID,
		AfterID:    afterID,
	})
	if err != nil {
		return nil, err
	}

	// TODO: remove custom Accept header when this API fully launch.
	req.Header.Set("Accept", mediaTypeReactionsPreview)

	return s.client.Do(ctx, req, nil)
}

// Get sub issues in an issue.
//
// GitHub API docs: https://docs.github.com/en/rest/issues/sub-issues?apiVersion=2022-11-28#list-sub-issues
//
//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/sub_issues
func (s *SubIssuesService) Get(ctx context.Context, owner string, repo string, issueNumber int) ([]*SubIssue, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues", owner, repo, issueNumber)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept header when this API fully launch.
	req.Header.Set("Accept", mediaTypeReactionsPreview)

	subIssues := new([]*SubIssue)
	resp, err := s.client.Do(ctx, req, subIssues)
	if err != nil {
		return nil, resp, err
	}

	return *subIssues, resp, nil
}
