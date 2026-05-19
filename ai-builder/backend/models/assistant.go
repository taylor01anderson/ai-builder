package models

import "gorm.io/gorm"

type Assistant struct {
    gorm.Model

    Name string
    Domain string
    Personality string
    SystemPrompt string

    UserID uint
}