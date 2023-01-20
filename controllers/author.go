package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/service"
)

type AuthorAPI interface {
	GetAuthorByID(c *gin.Context)
	GetAllAuthor(c *gin.Context)
	CreateNewAuthor(c *gin.Context)
	UpdateAuthor(c *gin.Context)
	DeleteAuthor(c *gin.Context)
}

type authorAPI struct {
	authorService service.AuthorService
}

func NewAuthorAPI(authorService service.AuthorService) *authorAPI {
	return &authorAPI{authorService}
}

func (a *authorAPI) GetAuthorByID(c *gin.Context) {
	idAuthor := c.Query("id")
	idAuthorInt, _ := strconv.Atoi(idAuthor)

	dataAuthor, err := a.authorService.GetAuthorByID(c.Request.Context(), idAuthorInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to get data author",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dataAuthor)
}

func (a *authorAPI) GetAllAuthor(c *gin.Context) {
	listDataAuthor, err := a.authorService.GetAllAuthor(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to get all data author",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, listDataAuthor)
}

func (a *authorAPI) CreateNewAuthor(c *gin.Context) {
	var input model.Author

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	newAuthor, err := a.authorService.CreateNewAuthor(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to create new author",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_author":   newAuthor.ID,
		"nama_author": newAuthor.NamaPengarang,
		"message":     "success create new author",
	})
}

func (a *authorAPI) UpdateAuthor(c *gin.Context) {
	var input model.Author

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	err = a.authorService.UpdateAuthor(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to update author",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update author",
	})
}

func (a *authorAPI) DeleteAuthor(c *gin.Context) {
	idAuthor := c.Query("id")
	idAuthorInt, _ := strconv.Atoi(idAuthor)

	err := a.authorService.DeleteAuthor(c.Request.Context(), idAuthorInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:       "failed to delete author",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete author",
	})
}
