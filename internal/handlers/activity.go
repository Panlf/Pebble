package handlers

import (
	"fmt"

	"Pebble/internal/database"
)

type ActivityHandler struct{}

func NewActivityHandler() *ActivityHandler {
	return &ActivityHandler{}
}

func (h *ActivityHandler) LogActivity(action, entityType, entityID, details string) error {
	_, err := database.DB.Exec(
		`INSERT INTO activity_logs (action, entity_type, entity_id, details) VALUES (?, ?, ?, ?)`,
		action, entityType, entityID, details,
	)
	if err != nil {
		return fmt.Errorf("failed to log activity: %w", err)
	}
	return nil
}

func (h *ActivityHandler) GetRecentActivities(limit int) ([]database.ActivityLog, error) {
	if limit <= 0 {
		limit = 50
	}

	rows, err := database.DB.Query(
		`SELECT id, action, entity_type, entity_id, details, created_at 
		 FROM activity_logs ORDER BY created_at DESC LIMIT ?`, limit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get activities: %w", err)
	}
	defer rows.Close()

	var activities []database.ActivityLog
	for rows.Next() {
		var a database.ActivityLog
		err := rows.Scan(&a.ID, &a.Action, &a.EntityType, &a.EntityID, &a.Details, &a.CreatedAt)
		if err != nil {
			continue
		}
		activities = append(activities, a)
	}
	return activities, nil
}

func (h *ActivityHandler) GetStatistics() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Project count
	var projectCount int
	err := database.DB.QueryRow(`SELECT COUNT(*) FROM projects`).Scan(&projectCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count projects: %w", err)
	}
	stats["projects"] = projectCount

	// Issue count
	var issueCount int
	err = database.DB.QueryRow(`SELECT COUNT(*) FROM issues`).Scan(&issueCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count issues: %w", err)
	}
	stats["issues"] = issueCount

	// Document count
	var docCount int
	err = database.DB.QueryRow(`SELECT COUNT(*) FROM documents`).Scan(&docCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count documents: %w", err)
	}
	stats["documents"] = docCount

	// Tag count
	var tagCount int
	err = database.DB.QueryRow(`SELECT COUNT(*) FROM tags`).Scan(&tagCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count tags: %w", err)
	}
	stats["tags"] = tagCount

	// Issue status distribution
	statusRows, err := database.DB.Query(
		`SELECT status, COUNT(*) FROM issues GROUP BY status`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get issue status: %w", err)
	}
	defer statusRows.Close()

	statusDist := make(map[string]int)
	for statusRows.Next() {
		var status string
		var count int
		if err := statusRows.Scan(&status, &count); err != nil {
			continue
		}
		statusDist[status] = count
	}
	stats["issue_status"] = statusDist

	// Recent activities count (last 7 days)
	var recentActivities int
	err = database.DB.QueryRow(
		`SELECT COUNT(*) FROM activity_logs WHERE created_at >= datetime('now', '-7 days')`,
	).Scan(&recentActivities)
	if err != nil {
		return nil, fmt.Errorf("failed to count recent activities: %w", err)
	}
	stats["recent_activities"] = recentActivities

	return stats, nil
}
