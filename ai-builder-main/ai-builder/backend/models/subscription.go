package models

import "gorm.io/gorm"

type Subscription struct {
    gorm.Model

    CustomerID string
    SubscriptionID string
    Plan string
    Status string

    UserID uint
}