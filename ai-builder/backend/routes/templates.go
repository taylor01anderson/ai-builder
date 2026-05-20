package routes

import "github.com/gin-gonic/gin"

func TemplateRoutes(rg *gin.RouterGroup) {

    rg.GET("/templates", GetTemplates)
}

func GetTemplates(c *gin.Context) {

    templates := []map[string]string{
        {
            "name": "Fitness Coach",
            "prompt": "You are an elite fitness coach.",
        },
        {
            "name": "Business Mentor",
            "prompt": "You help startups scale.",
        },
        {
            "name": "Coding Assistant",
            "prompt": "You are a senior engineer.",
        },
    }

    c.JSON(200, templates)
}