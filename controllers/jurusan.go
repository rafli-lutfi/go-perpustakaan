package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/service"
)

type JurusanAPI interface {
	GetAllJurusan(c *gin.Context)
}

type jurusanAPI struct {
	jurusanService service.JurusanService
}

func NewJurusanAPI(jurusanService service.JurusanService) *jurusanAPI {
	return &jurusanAPI{jurusanService}
}

func (j *jurusanAPI) GetAllJurusan(c *gin.Context) {
	_, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	allJurusan, err := j.jurusanService.GetAllJurusan(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "error internal server",
		})
		return
	}

	c.JSON(http.StatusOK, allJurusan)
}
