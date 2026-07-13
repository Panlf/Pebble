package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"project-document-system/internal/crypto"
	"project-document-system/internal/database"

	"github.com/google/uuid"
)

type DocumentHandler struct{}

func NewDocumentHandler() *DocumentHandler {
	return &DocumentHandler{}
}

func (h *DocumentHandler) UploadDocument(projectID string, issueID *string, title, fileName string, fileContent []byte, password string) (*database.Document, error) {
	project, err := (&ProjectHandler{}).GetProject(projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	salt, err := crypto.GenerateSalt()
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	encryptedName := crypto.EncryptFileName(fileName, password, salt)
	docID := uuid.New().String()

	storagePath := project.StoragePath
	if issueID != nil {
		storagePath = filepath.Join(storagePath, *issueID)
	} else {
		storagePath = filepath.Join(storagePath, "general")
	}

	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	filePath := filepath.Join(storagePath, encryptedName)
	if err := os.WriteFile(filePath, fileContent, 0644); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	fileType := filepath.Ext(fileName)
	fileSize := int64(len(fileContent))

	_, err = database.DB.Exec(
		`INSERT INTO documents (id, issue_id, project_id, title, file_path, encrypted_name, file_type, file_size) 
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		docID, issueID, projectID, title, filePath, encryptedName, fileType, fileSize,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert document: %w", err)
	}

	if issueID != nil {
		_, err = database.DB.Exec(
			`INSERT INTO issue_documents (issue_id, document_id) VALUES (?, ?)`,
			*issueID, docID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to link document to issue: %w", err)
		}
	}

	return h.GetDocument(docID)
}

func (h *DocumentHandler) GetDocument(id string) (*database.Document, error) {
	doc := &database.Document{}
	err := database.DB.QueryRow(
		`SELECT id, issue_id, project_id, title, file_path, encrypted_name, file_type, file_size, created_at, updated_at 
		 FROM documents WHERE id = ?`, id,
	).Scan(&doc.ID, &doc.IssueID, &doc.ProjectID, &doc.Title, &doc.FilePath,
		&doc.EncryptedName, &doc.FileType, &doc.FileSize, &doc.CreatedAt, &doc.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %w", err)
	}
	return doc, nil
}

func (h *DocumentHandler) ListDocuments(projectID string, issueID *string) ([]database.Document, error) {
	var rows *database.sql.Rows
	var err error

	if issueID != nil {
		rows, err = database.DB.Query(
			`SELECT d.id, d.issue_id, d.project_id, d.title, d.file_path, d.encrypted_name, d.file_type, d.file_size, d.created_at, d.updated_at 
			 FROM documents d
			 JOIN issue_documents id ON d.id = id.document_id
			 WHERE id.issue_id = ? ORDER BY d.created_at DESC`, *issueID,
		)
	} else {
		rows, err = database.DB.Query(
			`SELECT id, issue_id, project_id, title, file_path, encrypted_name, file_type, file_size, created_at, updated_at 
			 FROM documents WHERE project_id = ? ORDER BY created_at DESC`, projectID,
		)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to list documents: %w", err)
	}
	defer rows.Close()

	var docs []database.Document
	for rows.Next() {
		var d database.Document
		err := rows.Scan(&d.ID, &d.IssueID, &d.ProjectID, &d.Title, &d.FilePath,
			&d.EncryptedName, &d.FileType, &d.FileSize, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan document: %w", err)
		}
		docs = append(docs, d)
	}
	return docs, nil
}

func (h *DocumentHandler) DeleteDocument(id string) error {
	doc, err := h.GetDocument(id)
	if err != nil {
		return err
	}

	if err := os.Remove(doc.FilePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	_, err = database.DB.Exec(`DELETE FROM documents WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}
	return nil
}

func (h *DocumentHandler) ExportDocument(id string, outputPath string, password string) error {
	doc, err := h.GetDocument(id)
	if err != nil {
		return err
	}

	return crypto.DecryptFile(doc.FilePath, outputPath, password, "")
}
