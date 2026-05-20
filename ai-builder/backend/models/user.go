package models

import "gorm.io/gorm"

type User struct {
    gorm.Model

    Email    string `json:"email"`
    Password string `json:"password"`
}

/* 

Add after adding user authentication and associating assistants with users:

type User struct {
    gorm.Model

    Email string `gorm:"unique"`
    Password string

    Assistants []Assistant
}

*/