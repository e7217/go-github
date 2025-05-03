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

// AddSubIssueRequest, ReprioritizeSubIssueRequest 등 요청 구조체 정의
type AddSubIssueRequest struct {
	SubIssueID int `json:"sub_issue_id"`
}

type ReprioritizeSubIssueRequest struct {
	SubIssueID int `json:"sub_issue_id"`
	AfterID    int `json:"after_id"`
}

// Add, Remove, List, Reprioritize 메서드 정의 예시
func (s *SubIssuesService) Add(ctx context.Context, owner, repo string, issueNumber int, req *AddSubIssueRequest) (*Issue, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues", owner, repo, issueNumber)
	// 실제 요청 코드 작성
	return nil, nil, nil
}

func (s *SubIssuesService) Remove(ctx context.Context, owner, repo string, issueNumber, subIssueID int) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues/%d", owner, repo, issueNumber, subIssueID)
	// 실제 요청 코드 작성
	return nil, nil
}

func (s *SubIssuesService) List(ctx context.Context, owner, repo string, issueNumber int) ([]*Issue, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues", owner, repo, issueNumber)
	// 실제 요청 코드 작성
	return nil, nil, nil
}

func (s *SubIssuesService) Reprioritize(ctx context.Context, owner, repo string, issueNumber int, req *ReprioritizeSubIssueRequest) (*Issue, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/issues/%d/sub_issues/priority", owner, repo, issueNumber)
	// 실제 요청 코드 작성
	return nil, nil, nil
}
