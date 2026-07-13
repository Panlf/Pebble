package handlers

import (
	"fmt"

	"project-document-system/internal/database"

	"github.com/google/uuid"
)

type TagHandler struct{}

func NewTagHandler() *TagHandler {
	return &TagHandler{}
}

func (h *TagHandler) CreateTag(name, color string) (*database.Tag, error) {
	tagID := uuid.New().String()

	_, err := database.DB.Exec(
		`INSERT INTO tags (id, name, color) VALUES (?, ?, ?)`,
		tagID, name, color,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tag: %w", err)
	}

	return h.GetTag(tagID)
}

func (h *TagHandler) GetTag(id string) (*database.Tag, error) {
	tag := &database.Tag{}
	err := database.DB.QueryRow(
		`SELECT id, name, color, created_at FROM tags WHERE id = ?`, id,
	).Scan(&tag.ID, &tag.Name, &tag.Color, &tag.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag: %w", err)
	}
	return tag, nil
}

func (h *TagHandler) ListTags() ([]database.Tag, error) {
	rows, err := database.DB.Query(
		`SELECT id, name, color, created_at FROM tags ORDER BY name`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list tags: %w", err)
	}
	defer rows.Close()

	var tags []database.Tag
	for rows.Next() {
		var t database.Tag
		err := rows.Scan(&t.ID, &t.Name, &t.Color, &t.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag: %w", err)
		}
		tags = append(tags, t)
	}
	return tags, nil
}

func (h *TagHandler) UpdateTag(id, name, color string) error {
	_, err := database.DB.Exec(
		`UPDATE tags SET name = ?, color = ? WHERE id = ?`,
		name, color, id,
	)
	if err != nil {
		return fmt.Errorf("failed to update tag: %w", err)
	}
	return nil
}

func (h *TagHandler) DeleteTag(id string) error {
	_, err := database.DB.Exec(`DELETE FROM tags WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}
	return nil
}

func (h *TagHandler) AddTagToDocument(documentID, tagID string) error {
	_, err := database.DB.Exec(
		`INSERT INTO document_tags (document_id, tag_id) VALUES (?, ?)`,
		documentID, tagID,
	)
	if err != nil {
		return fmt.Errorf("failed to add tag to document: %w", err)
	}
	return nil
}

func (h *TagHandler) RemoveTagFromDocument(documentID, tagID string) error {
	_, err := database.DB.Exec(
		`DELETE FROM document_tags WHERE document_id = ? AND tag_id = ?`,
		documentID, tagID,
	)
	if err != nil {
		return fmt.Errorf("failed to remove tag from document: %w", err)
	}
	return nil
}

func (h *TagHandler) GetDocumentTags(documentID string) ([]database.Tag, error) {
	rows, err := database.DB.Query(
		`SELECT t.id, t.name, t.color, t.created_at 
		 FROM tags t
		 JOIN document_tags dt ON t.id = dt.tag_id
		 WHERE dt.document_id = ? ORDER BY t.name`, documentID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get document tags: %w", err)
	}
	defer rows.Close()

	var tags []database.Tag
	for rows.Next() {
		var t database.Tag
		err := rows.Scan(&t.ID, &t.Name, &t.Color, &t.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tag: %w", err)
		}
		tags = append(tags, t)
	}
	return tags, nil
}
