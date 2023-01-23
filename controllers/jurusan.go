package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/service"
)

type JurusanAPI interface {
	GetAllJurusan(c *gin.Context)
	CreateJurusan(c *gin.Context)
	UpdateJurusan(c *gin.Context)
	DeleteJurusan(c *gin.Context)
}

type jurusanAPI struct {
	jurusanService service.JurusanService
}

func NewJurusanAPI(jurusanService service.JurusanService) *jurusanAPI {
	return &jurusanAPI{jurusanService}
}

func (j *jurusanAPI) GetAllJurusan(c *gin.Context) {
	// Get offset and limit page
	path := c.Request.URL.Path
	offsetString := c.DefaultQuery("offset", "0")
	limitString := c.DefaultQuery("limit", "20")

	// convert to integer
	offset, _ := strconv.Atoi(offsetString)
	limit, _ := strconv.Atoi(limitString)

	allJurusan, count, err := j.jurusanService.GetAllJurusan(c.Request.Context(), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "error internal server",
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
			Result:   allJurusan,
		},
		"message": "success get all jurusan",
	})
}

func (j *jurusanAPI) CreateJurusan(c *gin.Context) {
	var input model.Jurusan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	if input.NamaJurusan == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "form data is empty",
			Description: "please check your form before submit",
		})
		return
	}

	newJurusan, err := j.jurusanService.NewJurusan(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to create new jurusan",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id_jurusan":   newJurusan.ID,
			"nama_jurusan": newJurusan.NamaJurusan,
		},
		"message": "success create new jurusan",
	})
}

func (j *jurusanAPI) UpdateJurusan(c *gin.Context) {
	var input model.Jurusan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	if input.NamaJurusan == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "form data is empty",
			Description: "please check your form before submit",
		})
		return
	}

	_, err = j.jurusanService.UpdateJurusan(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to update jurusan",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success update jurusan",
	})
}

func (j *jurusanAPI) DeleteJurusan(c *gin.Context) {
	var input model.Jurusan

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	if input.ID <= 0 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "form data is empty",
			Description: "please check your form before submit",
		})
		return
	}

	err = j.jurusanService.DeleteJurusan(c.Request.Context(), input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to delete jurusan",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success delete jurusan",
	})
}
