package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/service"
)

type BukuAPI interface {
	GetBukuByID(c *gin.Context)
	GetAllBuku(c *gin.Context)
	CreateBuku(c *gin.Context)
	UpdateBuku(c *gin.Context)
	DeleteBuku(c *gin.Context)
}

type bukuAPI struct {
	bukuService service.BukuService
}

func NewBukuAPI(bukuService service.BukuService) *bukuAPI {
	return &bukuAPI{bukuService}
}

func (b *bukuAPI) GetBukuByID(c *gin.Context) {
	idBuku := c.Query("id")
	idBukuInt, _ := strconv.Atoi(idBuku)

	dataBuku, err := b.bukuService.GetBukuByID(c.Request.Context(), idBukuInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to get data buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dataBuku)
}

func (b *bukuAPI) GetAllBuku(c *gin.Context) {
	listDataBuku, err := b.bukuService.GetAllBuku(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to get all data buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, listDataBuku)
}

func (b *bukuAPI) CreateBuku(c *gin.Context) {
	input := model.Buku{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	newBuku, err := b.bukuService.CreateNewBuku(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to create new buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_buku":    newBuku.ID,
		"judul_buku": newBuku.JudulBuku,
		"message":    "success create new buku",
	})
}

func (b *bukuAPI) UpdateBuku(c *gin.Context) {
	input := model.Buku{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	err = b.bukuService.UpdateBuku(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to update buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update buku",
	})
}

func (b *bukuAPI) DeleteBuku(c *gin.Context) {
	idBuku := c.Query("id_buku")
	idBukuInt, _ := strconv.Atoi(idBuku)

	err := b.bukuService.DeleteBuku(c.Request.Context(), idBukuInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to delete buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete buku",
	})
}
