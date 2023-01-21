package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/service"
	"golang.org/x/crypto/bcrypt"
)

type UserAPI interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
	Delete(c *gin.Context)
	GetUserData(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Login(c *gin.Context) {
	var user model.UserLogin

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "email or password is empty",
			Description: "email or password is empty",
		})
		return
	}

	userData := model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	id, err := u.userService.Login(c.Request.Context(), &userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to login",
			Description: err.Error(),
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": fmt.Sprint(id),
		"expiry":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "invalid to create token",
			Description: err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("session_token", tokenString, 3600*24*2, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"user_id": id,
		},
		"message": "login success",
	})
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" || user.NPM == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "register data is empty",
			Description: "please check your data before register",
		})
		return
	}

	// Hash The Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to hash password",
			Description: err.Error(),
		})

		return
	}

	userData := model.User{
		Fullname:    user.Fullname,
		Email:       user.Email,
		Password:    string(hash),
		NPM:         user.NPM,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		IDJurusan:   user.IDJurusan,
	}

	newUser, err := u.userService.Register(c.Request.Context(), &userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to create user",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"user_id": newUser.ID,
		},
		"message": "register success",
	})
}

func (u *userAPI) Logout(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "logout sucess",
	})
}

func (u *userAPI) Delete(c *gin.Context) {
	userID := c.Query("user_id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  "user id is empty",
		})
		return
	}

	userIDInt, _ := strconv.Atoi(userID)

	err := u.userService.Delete(c.Request.Context(), userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error internal server",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success delete",
	})
}

func (u *userAPI) GetUserData(c *gin.Context) {
	userID, _ := c.Get("id")

	userIDInt := userID.(int)

	userInfo, err := u.userService.GetUserData(c.Request.Context(), userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusInternalServerError,
			Error:       "failed to get user data",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    userInfo,
		"message": "success get user data",
	})
}

func (u *userAPI) UpdateUser(c *gin.Context) {
	var input model.User

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to read body",
			Description: err.Error(),
		})
		return
	}

	err = u.userService.Update(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Status:      http.StatusBadRequest,
			Error:       "failed to update user",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success delete user",
	})
}
