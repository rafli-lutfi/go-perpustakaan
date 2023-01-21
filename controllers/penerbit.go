package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/service"
)

type PenerbitAPI interface {
	GetPenerbitByID(c *gin.Context)
	GetAllPenerbit(c *gin.Context)
	CreateNewPenerbit(c *gin.Context)
	UpdatePenerbit(c *gin.Context)
	DeletePenerbit(c *gin.Context)
}

type penerbitAPI struct {
	penerbitService service.PenerbitService
}

func NewPenerbitAPI(penerbitService service.PenerbitService) *penerbitAPI {
	return &penerbitAPI{penerbitService}
}

func (p *penerbitAPI) GetPenerbitByID(c *gin.Context) {
	idPenerbit := c.Query("id")
	idPenerbitInt, _ := strconv.Atoi(idPenerbit)

	dataPenerbit, err := p.penerbitService.GetPenerbitByID(c.Request.Context(), idPenerbitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to get data penerbit",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    dataPenerbit,
		"message": "success get penerbit",
	})
}

func (p *penerbitAPI) GetAllPenerbit(c *gin.Context) {
	listDataPenerbit, err := p.penerbitService.GetAllPenerbit(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to get all data penerbit",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    listDataPenerbit,
		"message": "success get all penerbit",
	})
}

func (p *penerbitAPI) CreateNewPenerbit(c *gin.Context) {
	var input model.Penerbit

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	newPenerbit, err := p.penerbitService.CreateNewPenerbit(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to create new penerbit",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id_penerbit":   newPenerbit.ID,
			"nama_penerbit": newPenerbit.NamaPenerbit,
		},
		"message": "success create new penerbit",
	})
}

func (p *penerbitAPI) UpdatePenerbit(c *gin.Context) {
	var input model.Penerbit

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	err = p.penerbitService.UpdatePenerbit(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to update penerbit",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success update penerbit",
	})
}

func (p *penerbitAPI) DeletePenerbit(c *gin.Context) {
	idPenerbit := c.Query("id")
	idPenerbitInt, _ := strconv.Atoi(idPenerbit)

	err := p.penerbitService.DeletePenerbit(c.Request.Context(), idPenerbitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to delete penerbit",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success delete penerbit",
	})
}
