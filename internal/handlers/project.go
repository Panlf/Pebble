package handlers

import (
	"fmt"
	"time"

	"project-document-system/internal/crypto"
	"project-document-system/internal/database"

	"github.com/google/uuid"
)

type ProjectHandler struct{}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

func (h *ProjectHandler) CreateProject(name, description, password string) (*database.Project, error) {
	salt, err := crypto.GenerateSalt()
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	passwordHash := crypto.HashPassword(password, salt)
	projectID := uuid.New().String()

	_, err = database.DB.Exec(
		`INSERT INTO projects (id, name, description, password_hash, salt, storage_path) 
		 VALUES (?, ?, ?, ?, ?, ?)`,
		projectID, name, description, passwordHash, salt, crypto.GetStoragePath(projectID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	return h.GetProject(projectID)
}

func (h *ProjectHandler) GetProject(id string) (*database.Project, error) {
	project := &database.Project{}
	err := database.DB.QueryRow(
		`SELECT id, name, description, password_hash, salt, created_at, updated_at, storage_path 
		 FROM projects WHERE id = ?`, id,
	).Scan(&project.ID, &project.Name, &project.Description, &project.PasswordHash,
		&project.Salt, &project.CreatedAt, &project.UpdatedAt, &project.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	return project, nil
}

func (h *ProjectHandler) ListProjects() ([]database.Project, error) {
	rows, err := database.DB.Query(
		`SELECT id, name, description, created_at, updated_at, storage_path 
		 FROM projects ORDER BY updated_at DESC`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	defer rows.Close()

	var projects []database.Project
	for rows.Next() {
		var p database.Project
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt, &p.StoragePath)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (h *ProjectHandler) UpdateProject(id, name, description string) error {
	_, err := database.DB.Exec(
		`UPDATE projects SET name = ?, description = ?, updated_at = ? WHERE id = ?`,
		name, description, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}
	return nil
}

func (h *ProjectHandler) DeleteProject(id string) error {
	_, err := database.DB.Exec(`DELETE FROM projects WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}

func (h *ProjectHandler) VerifyPassword(id, password string) (bool, error) {
	project, err := h.GetProject(id)
	if err != nil {
		return false, err
	}

	hash := crypto.HashPassword(password, project.Salt)
	return hash == project.PasswordHash, nil
}

func (h *ProjectHandler) ChangePassword(id, oldPassword, newPassword string) error {
	project, err := h.GetProject(id)
	if err != nil {
		return err
	}

	oldHash := crypto.HashPassword(oldPassword, project.Salt)
	if oldHash != project.PasswordHash {
		return fmt.Errorf("incorrect old password")
	}

	salt, err := crypto.GenerateSalt()
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	newHash := crypto.HashPassword(newPassword, salt)
	_, err = database.DB.Exec(
		`UPDATE projects SET password_hash = ?, salt = ?, updated_at = ? WHERE id = ?`,
		newHash, salt, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("failed to change password: %w", err)
	}
	return nil
}
