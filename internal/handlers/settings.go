package handlers

import (
	"fmt"

	"Pebble/internal/database"
)

type SettingsHandler struct{}

func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{}
}

func (h *SettingsHandler) GetSetting(key string) (string, error) {
	var value string
	err := database.DB.QueryRow(`SELECT value FROM settings WHERE key = ?`, key).Scan(&value)
	if err != nil {
		return "", fmt.Errorf("failed to get setting: %w", err)
	}
	return value, nil
}

func (h *SettingsHandler) SetSetting(key, value string) error {
	_, err := database.DB.Exec(
		`INSERT INTO settings (key, value) VALUES (?, ?) 
		 ON CONFLICT(key) DO UPDATE SET value = excluded.value, updated_at = CURRENT_TIMESTAMP`,
		key, value,
	)
	if err != nil {
		return fmt.Errorf("failed to set setting: %w", err)
	}
	return nil
}

func (h *SettingsHandler) GetAllSettings() (map[string]string, error) {
	rows, err := database.DB.Query(`SELECT key, value FROM settings`)
	if err != nil {
		return nil, fmt.Errorf("failed to get settings: %w", err)
	}
	defer rows.Close()

	settings := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			continue
		}
		settings[key] = value
	}
	return settings, nil
}
