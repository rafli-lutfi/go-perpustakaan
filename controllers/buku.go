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
			Status:      http.StatusInternalServerError,
			Error:       "failed to get data buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    dataBuku,
		"message": "success get buku",
	})
}

func (b *bukuAPI) GetAllBuku(c *gin.Context) {
	// Get offset and limit page
	path := c.Request.URL.Path
	offsetString := c.DefaultQuery("offset", "0")
	limitString := c.DefaultQuery("limit", "20")

	// convert to integer
	offset, _ := strconv.Atoi(offsetString)
	limit, _ := strconv.Atoi(limitString)

	listDataBuku, count, err := b.bukuService.GetAllBuku(c.Request.Context(), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to get all data buku",
			Description: err.Error(),
		})
		return
	}

	var previousPage string
	var nextPage string

	if offset+limit <= count {
		newOffset := offset + limit
		newOffsetString := strconv.Itoa(newOffset)
		nextPage = path + "/?offset=" + newOffsetString + "&limit=" + limitString
	}

	if offset > 0 {
		newOffset := offset - limit
		newOffsetString := strconv.Itoa(newOffset)
		previousPage = path + "/?offset=" + newOffsetString + "&limit=" + limitString
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": model.Pagination{
			Count:    count,
			Next:     nextPage,
			Previous: previousPage,
			Result:   listDataBuku,
		},
		"message": "success get all buku",
	})
}

func (b *bukuAPI) CreateBuku(c *gin.Context) {
	input := model.Buku{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	newBuku, err := b.bukuService.CreateNewBuku(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to create new buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id_buku":    newBuku.ID,
			"judul_buku": newBuku.JudulBuku,
		},
		"message": "success create new buku",
	})
}

func (b *bukuAPI) UpdateBuku(c *gin.Context) {
	input := model.Buku{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	err = b.bukuService.UpdateBuku(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to update buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success update buku",
	})
}

func (b *bukuAPI) DeleteBuku(c *gin.Context) {
	idBuku := c.Query("id_buku")
	idBukuInt, _ := strconv.Atoi(idBuku)

	err := b.bukuService.DeleteBuku(c.Request.Context(), idBukuInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to delete buku",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success delete buku",
	})
}
