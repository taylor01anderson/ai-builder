package models

import "gorm.io/gorm"

type Assistant struct {
	gorm.Model

	Name         string `json:"name"`
	Domain       string `json:"domain"`
	Personality  string `json:"personality"`
	SystemPrompt string `json:"prompt"`

    // Add back UserID if you want to associate assistants with users
    // UserID uint
}

