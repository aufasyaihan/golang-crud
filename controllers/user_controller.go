package controllers

import (
    "golang-crud-krs/config"
    "golang-crud-krs/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateKRS(c *gin.Context) {
    var krs models.KRS
    if err := c.ShouldBindJSON(&krs); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&krs)
    c.JSON(http.StatusCreated, krs)
}

func GetAllKRS(c *gin.Context) {
    var allKRS []models.KRS
    config.DB.Order("id_krs ASC").Find(&allKRS)
    c.JSON(http.StatusOK, allKRS)
}

func GetKRS(c *gin.Context) {
    var krs models.KRS
    if err := config.DB.First(&krs, c.Param("id_krs")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "KRS not found"})
        return
    }
    c.JSON(http.StatusOK, krs)
}

func UpdateKRS(c *gin.Context) {
    var krs models.KRS
    if err := config.DB.First(&krs, c.Param("id_krs")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "KRS not found"})
        return
    }
    if err := c.ShouldBindJSON(&krs); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&krs)
    c.JSON(http.StatusOK, krs)
}

func DeleteKRS(c *gin.Context) {
    var krs models.KRS
    if err := config.DB.First(&krs, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "KRS not found"})
        return
    }
    config.DB.Delete(&krs)
    c.JSON(http.StatusOK, gin.H{"message": "KRS deleted successfully"})
}
