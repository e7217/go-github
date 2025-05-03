package github

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSubIssueService_Add(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/issues/1/sub_issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		// 요청 body, 헤더 등 검증 필요시 추가
		fmt.Fprint(w, `{"id": 2, "number": 2, "title": "sub issue"}`)
	})

	input := &AddSubIssueRequest{
		SubIssueID: 2,
	}
	ctx := context.Background()
	subIssue, _, err := client.SubIssues.Add(ctx, "o", "r", 1, input)
	if err != nil {
		t.Errorf("SubIssues.Add returned error: %v", err)
	}

	want := &Issue{ID: Ptr(int64(2)), Number: Ptr(2), Title: Ptr("sub issue")}
	if !cmp.Equal(subIssue, want) {
		t.Errorf("SubIssues.Add returned %+v, want %+v", subIssue, want)
	}
}

func TestSubIssueService_Remove(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/issues/1/sub_issues/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.SubIssues.Remove(ctx, "o", "r", 1, 2)
	if err != nil {
		t.Errorf("SubIssues.Remove returned error: %v", err)
	}
}

func TestSubIssueService_List(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/issues/1/sub_issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id": 2, "number": 2, "title": "sub issue"}]`)
	})

	ctx := context.Background()
	subIssues, _, err := client.SubIssues.List(ctx, "o", "r", 1)
	if err != nil {
		t.Errorf("SubIssues.List returned error: %v", err)
	}

	want := []*Issue{{ID: Ptr(int64(2)), Number: Ptr(2), Title: Ptr("sub issue")}}
	if !cmp.Equal(subIssues, want) {
		t.Errorf("SubIssues.List returned %+v, want %+v", subIssues, want)
	}
}

func TestSubIssueService_Reprioritize(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)

	mux.HandleFunc("/repos/o/r/issues/1/sub_issues/priority", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id": 2, "number": 2, "title": "sub issue"}`)
	})

	input := &ReprioritizeSubIssueRequest{
		SubIssueID: 2,
		AfterID:    1,
	}
	ctx := context.Background()
	subIssue, _, err := client.SubIssues.Reprioritize(ctx, "o", "r", 1, input)
	if err != nil {
		t.Errorf("SubIssues.Reprioritize returned error: %v", err)
	}

	want := &Issue{ID: Ptr(int64(2)), Number: Ptr(2), Title: Ptr("sub issue")}
	if !cmp.Equal(subIssue, want) {
		t.Errorf("SubIssues.Reprioritize returned %+v, want %+v", subIssue, want)
	}
}
