package routes

import (
    "ai-builder/config"
    "ai-builder/models"
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func AuthRoutes(rg *gin.RouterGroup) {

    rg.POST("/register", Register)
}

func Register(c *gin.Context) {

    var body struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }

    c.BindJSON(&body)

    hashed, _ := bcrypt.GenerateFromPassword(
        []byte(body.Password),
        14,
    )

    user := models.User{
        Email: body.Email,
        Password: string(hashed),
    }

    config.DB.Create(&user)

    c.JSON(http.StatusOK, gin.H{
        "message": "registered",
    })
}