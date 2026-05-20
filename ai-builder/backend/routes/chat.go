package routes

import (
	"ai-builder/config"
	"ai-builder/models"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(rg *gin.RouterGroup) {

    rg.POST("/chat", Chat)
}

type ChatRequest struct {
	Prompt      string `json:"prompt"`
	AssistantID uint   `json:"assistant_id"`
}

func Chat(c *gin.Context) {

	var req ChatRequest

	if err := c.BindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	var assistant models.Assistant

	config.DB.First(
		&assistant,
		req.AssistantID,
	)
}