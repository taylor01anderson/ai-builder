package services

import (
    "ai-builder/config"
    "ai-builder/models"
    "strings"
)

func SearchDocuments(query string) []models.Document {

    var docs []models.Document

    config.DB.Find(&docs)

    var relevant []models.Document

    for _, doc := range docs {

        if strings.Contains(
            strings.ToLower(doc.Content),
            strings.ToLower(query),
        ) {
            relevant = append(relevant, doc)
        }
    }

    return relevant
}