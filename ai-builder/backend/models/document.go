package models

import "gorm.io/gorm"

type Document struct {
    gorm.Model

    FileName string
    Content string `gorm:"type:text"`

    AssistantID uint
}