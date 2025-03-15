package controllers

import (
	"golang-crud-krs/config"
	"golang-crud-krs/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KRSResponse struct {
	Nim          string `json:"nim"`
	KodeMatakuliah string `json:"kode_matakuliah"`
	Matakuliah   string `json:"matakuliah"`
	Semester     string `json:"semester"`
	TahunAkademik string `json:"tahunakademik"`
}

func convertToResponse(krs models.KRS) KRSResponse {
	return KRSResponse{
			Nim:           krs.Nim,
			KodeMatakuliah: krs.KodeMatakuliah,
			Matakuliah:    krs.Matakuliah,
			Semester:      krs.Semester,
			TahunAkademik: krs.TahunAkademik,
	}
}

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
	config.DB.Order("id_krs ASC").Limit(10).Find(&allKRS)

	var response []KRSResponse
	for _, krs := range allKRS {
			response = append(response, convertToResponse(krs))
	}

	c.JSON(http.StatusOK, response)
}

func GetKRS(c *gin.Context) {
	var krs models.KRS
	if err := config.DB.First(&krs, c.Param("id_krs")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "KRS not found"})
			return
	}
	c.JSON(http.StatusOK, convertToResponse(krs))
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
