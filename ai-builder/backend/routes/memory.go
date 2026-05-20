package routes

import (
    "ai-builder/config"
    "ai-builder/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func MemoryRoutes(rg *gin.RouterGroup) {

    rg.POST("/memory", SaveMessage)
    rg.GET("/memory/:assistantId", GetMemory)
}

func SaveMessage(c *gin.Context) {

    var body struct {
        Role string `json:"role"`
        Content string `json:"content"`
        AssistantID uint `json:"assistant_id"`
    }

    c.BindJSON(&body)

    message := models.Message{
        Role: body.Role,
        Content: body.Content,
        AssistantID: body.AssistantID,
    }

    config.DB.Create(&message)

    c.JSON(http.StatusOK, message)
}

func GetMemory(c *gin.Context) {

    assistantId := c.Param("assistantId")

    var messages []models.Message

    config.DB.Where(
        "assistant_id = ?",
        assistantId,
    ).Find(&messages)

    c.JSON(http.StatusOK, messages)
}