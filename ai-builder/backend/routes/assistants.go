package routes

import (
	"ai-builder/config"
	"ai-builder/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssistantRoutes(rg *gin.RouterGroup) {

    rg.POST("/assistants", CreateAssistant)
    rg.GET("/assistants", GetAssistants)
}

func CreateAssistant(c *gin.Context) {

    var body struct {
        Name string `json:"name"`
        Domain string `json:"domain"`
        Personality string `json:"personality"`
        SystemPrompt string `json:"system_prompt"`
    }

    c.BindJSON(&body)

    assistant := models.Assistant{
        Name: body.Name,
        Domain: body.Domain,
        Personality: body.Personality,
        SystemPrompt: body.SystemPrompt,
    }

    config.DB.Create(&assistant)

    c.JSON(http.StatusOK, assistant)
}

func GetAssistants(c *gin.Context) {

    var assistants []models.Assistant

    config.DB.Find(&assistants)

    c.JSON(http.StatusOK, assistants)
}