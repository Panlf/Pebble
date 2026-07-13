package handlers

import (
	"fmt"
	"time"

	"Pebble/internal/database"

	"github.com/google/uuid"
)

type IssueHandler struct{}

func NewIssueHandler() *IssueHandler {
	return &IssueHandler{}
}

func (h *IssueHandler) CreateIssue(projectID, title, description, contactPerson string) (*database.Issue, error) {
	issueID := uuid.New().String()

	_, err := database.DB.Exec(
		`INSERT INTO issues (id, project_id, title, description, contact_person) 
		 VALUES (?, ?, ?, ?, ?)`,
		issueID, projectID, title, description, contactPerson,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create issue: %w", err)
	}

	return h.GetIssue(issueID)
}

func (h *IssueHandler) GetIssue(id string) (*database.Issue, error) {
	issue := &database.Issue{}
	err := database.DB.QueryRow(
		`SELECT id, project_id, title, description, status, contact_person, created_at, updated_at 
		 FROM issues WHERE id = ?`, id,
	).Scan(&issue.ID, &issue.ProjectID, &issue.Title, &issue.Description,
		&issue.Status, &issue.ContactPerson, &issue.CreatedAt, &issue.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get issue: %w", err)
	}
	return issue, nil
}

func (h *IssueHandler) ListIssues(projectID string) ([]database.Issue, error) {
	rows, err := database.DB.Query(
		`SELECT id, project_id, title, description, status, contact_person, created_at, updated_at 
		 FROM issues WHERE project_id = ? ORDER BY created_at DESC`, projectID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list issues: %w", err)
	}
	defer rows.Close()

	var issues []database.Issue
	for rows.Next() {
		var i database.Issue
		err := rows.Scan(&i.ID, &i.ProjectID, &i.Title, &i.Description,
			&i.Status, &i.ContactPerson, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan issue: %w", err)
		}
		issues = append(issues, i)
	}
	return issues, nil
}

func (h *IssueHandler) UpdateIssue(id, title, description, status, contactPerson string) error {
	_, err := database.DB.Exec(
		`UPDATE issues SET title = ?, description = ?, status = ?, contact_person = ?, updated_at = ? WHERE id = ?`,
		title, description, status, contactPerson, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("failed to update issue: %w", err)
	}
	return nil
}

func (h *IssueHandler) DeleteIssue(id string) error {
	_, err := database.DB.Exec(`DELETE FROM issues WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete issue: %w", err)
	}
	return nil
}

func (h *IssueHandler) UpdateStatus(id, status string) error {
	_, err := database.DB.Exec(
		`UPDATE issues SET status = ?, updated_at = ? WHERE id = ?`,
		status, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}
	return nil
}
