package routes

import (
    "ai-builder/config"
    "ai-builder/models"

    "github.com/gin-gonic/gin"
)

func MarketplaceRoutes(rg *gin.RouterGroup) {

    rg.GET("/marketplace", Marketplace)
}

func Marketplace(c *gin.Context) {

    var assistants []models.Assistant

    config.DB.Find(&assistants)

    c.JSON(200, assistants)
}