package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/service"
)

type KategoriAPI interface {
	GetKategoriByID(c *gin.Context)
	GetAllKategori(c *gin.Context)
	CreateNewKategori(c *gin.Context)
	UpdateKategori(c *gin.Context)
	DeleteKategori(c *gin.Context)
}

type kategoriAPI struct {
	kategoriService service.KategoriService
}

func NewKategoriAPI(kategoriService service.KategoriService) *kategoriAPI {
	return &kategoriAPI{kategoriService}
}

func (k *kategoriAPI) GetKategoriByID(c *gin.Context) {
	idKategori := c.Query("id")
	idKategoriInt, _ := strconv.Atoi(idKategori)

	dataKategori, err := k.kategoriService.GetKategoriByID(c.Request.Context(), idKategoriInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to get data kategori",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    dataKategori,
		"message": "success get kategori",
	})
}

func (k *kategoriAPI) GetAllKategori(c *gin.Context) {
	listDataKategori, err := k.kategoriService.GetAllKategori(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to get all data kategori",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    listDataKategori,
		"message": "success get all kategori",
	})
}

func (k *kategoriAPI) CreateNewKategori(c *gin.Context) {
	var input model.Kategori

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	newKategori, err := k.kategoriService.CreateNewKategori(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to create new kategori",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id_kategori":   newKategori.ID,
			"nama_kategori": newKategori.NamaKategori,
		},
		"message": "success create new kategori",
	})
}

func (k *kategoriAPI) UpdateKategori(c *gin.Context) {
	var input model.Kategori

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	err = k.kategoriService.UpdateKategori(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to update kategori",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success update kategori",
	})
}

func (k *kategoriAPI) DeleteKategori(c *gin.Context) {
	idKategori := c.Query("id")
	idKategoriInt, _ := strconv.Atoi(idKategori)

	err := k.kategoriService.DeleteKategori(c.Request.Context(), idKategoriInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to delete kategori",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success delete kategori",
	})
}
