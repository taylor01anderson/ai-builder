package routes

import (
    "ai-builder/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

func ChatRoutes(rg *gin.RouterGroup) {

    rg.POST("/chat", Chat)
}

func Chat(c *gin.Context) {

    var body struct {
        Prompt string `json:"prompt"`
    }

    c.BindJSON(&body)

    response, err := services.GenerateResponse(body.Prompt)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })

        return
    }

    c.JSON(http.StatusOK, gin.H{
        "response": response,
    })
}