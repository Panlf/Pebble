package handlers

import (
	"fmt"

	"project-document-system/internal/database"
)

type SearchHandler struct{}

func NewSearchHandler() *SearchHandler {
	return &SearchHandler{}
}

type SearchResult struct {
	Type      string      `json:"type"`
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	Snippet   string      `json:"snippet"`
	ProjectID string      `json:"project_id"`
	CreatedAt string      `json:"created_at"`
}

func (h *SearchHandler) Search(query string, projectID string) ([]SearchResult, error) {
	if query == "" {
		return []SearchResult{}, nil
	}

	var results []SearchResult

	// Search projects
	projectRows, err := database.DB.Query(
		`SELECT id, name, description, created_at FROM projects 
		 WHERE name LIKE ? OR description LIKE ?`,
		"%"+query+"%", "%"+query+"%",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to search projects: %w", err)
	}
	defer projectRows.Close()

	for projectRows.Next() {
		var p database.Project
		err := projectRows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
		if err != nil {
			continue
		}
		results = append(results, SearchResult{
			Type:      "project",
			ID:        p.ID,
			Title:     p.Name,
			Snippet:   truncate(p.Description, 100),
			ProjectID: p.ID,
			CreatedAt: p.CreatedAt.Format("2006-01-02"),
		})
	}

	// Search issues
	issueQuery := `SELECT id, project_id, title, description, created_at FROM issues 
				   WHERE (title LIKE ? OR description LIKE ?)`
	issueArgs := []interface{}{"%" + query + "%", "%" + query + "%"}

	if projectID != "" {
		issueQuery += " AND project_id = ?"
		issueArgs = append(issueArgs, projectID)
	}

	issueRows, err := database.DB.Query(issueQuery, issueArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to search issues: %w", err)
	}
	defer issueRows.Close()

	for issueRows.Next() {
		var i database.Issue
		err := issueRows.Scan(&i.ID, &i.ProjectID, &i.Title, &i.Description, &i.CreatedAt)
		if err != nil {
			continue
		}
		results = append(results, SearchResult{
			Type:      "issue",
			ID:        i.ID,
			Title:     i.Title,
			Snippet:   truncate(i.Description, 100),
			ProjectID: i.ProjectID,
			CreatedAt: i.CreatedAt.Format("2006-01-02"),
		})
	}

	// Search documents
	docQuery := `SELECT id, project_id, title, created_at FROM documents 
				 WHERE title LIKE ?`
	docArgs := []interface{}{"%" + query + "%"}

	if projectID != "" {
		docQuery += " AND project_id = ?"
		docArgs = append(docArgs, projectID)
	}

	docRows, err := database.DB.Query(docQuery, docArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to search documents: %w", err)
	}
	defer docRows.Close()

	for docRows.Next() {
		var d database.Document
		err := docRows.Scan(&d.ID, &d.ProjectID, &d.Title, &d.CreatedAt)
		if err != nil {
			continue
		}
		results = append(results, SearchResult{
			Type:      "document",
			ID:        d.ID,
			Title:     d.Title,
			Snippet:   "文档",
			ProjectID: d.ProjectID,
			CreatedAt: d.CreatedAt.Format("2006-01-02"),
		})
	}

	// Search tags
	tagRows, err := database.DB.Query(
		`SELECT id, name, created_at FROM tags WHERE name LIKE ?`,
		"%"+query+"%",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to search tags: %w", err)
	}
	defer tagRows.Close()

	for tagRows.Next() {
		var t database.Tag
		err := tagRows.Scan(&t.ID, &t.Name, &t.CreatedAt)
		if err != nil {
			continue
		}
		results = append(results, SearchResult{
			Type:      "tag",
			ID:        t.ID,
			Title:     t.Name,
			Snippet:   "标签",
			ProjectID: "",
			CreatedAt: t.CreatedAt.Format("2006-01-02"),
		})
	}

	return results, nil
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
