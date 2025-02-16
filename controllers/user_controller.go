package controllers

import (
    "golang-crud-users/config"
    "golang-crud-users/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&user)
    c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
    var users []models.User
    config.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
    var user models.User
    if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    var user models.User
    if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&user)
    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    var user models.User
    if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    config.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
