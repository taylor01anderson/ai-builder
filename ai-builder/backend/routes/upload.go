package routes

import (
    "ai-builder/config"
    "ai-builder/models"
    "io"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

func UploadRoutes(rg *gin.RouterGroup) {

    rg.POST("/upload", UploadDocument)
}

func UploadDocument(c *gin.Context) {

    file, err := c.FormFile("file")

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })

        return
    }

    openedFile, _ := file.Open()

    defer openedFile.Close()

    bytes, _ := io.ReadAll(openedFile)

    document := models.Document{
        FileName: file.Filename,
        Content: string(bytes),
    }

    config.DB.Create(&document)

    os.WriteFile(
        "uploads/"+file.Filename,
        bytes,
        0644,
    )

    c.JSON(http.StatusOK, document)
}