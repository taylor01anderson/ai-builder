package models

import "gorm.io/gorm"

type Message struct {
    gorm.Model

    Role string
    Content string `gorm:"type:text"`

    AssistantID uint
}