package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/controllers"
	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/middleware"
	"github.com/rafli-lutfi/perpustakaan/service"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler    controllers.UserAPI
	JurusanAPIHandler controllers.JurusanAPI
}

func RunServer(db *gorm.DB, r *gin.Engine) *gin.Engine {
	userDatabase := database.NewUserDatabase(db)
	jurusanDatabase := database.NewJurusanDatabase(db)

	userService := service.NewUserService(userDatabase, jurusanDatabase)
	jurusanService := service.NewJurusanService(jurusanDatabase)

	userAPIHandler := controllers.NewUserAPI(userService)
	jurusanAPIHandler := controllers.NewJurusanAPI(jurusanService)

	apiHandler := APIHandler{
		UserAPIHandler:    userAPIHandler,
		JurusanAPIHandler: jurusanAPIHandler,
	}

	server := r.Group("/api/v1")
	users := server.Group("/users")

	users.POST("/login", apiHandler.UserAPIHandler.Login)
	users.POST("/register", apiHandler.UserAPIHandler.Register)
	users.POST("/logout", apiHandler.UserAPIHandler.Logout)
	users.DELETE("/delete", apiHandler.UserAPIHandler.Delete)
	users.GET("/info", middleware.Auth, apiHandler.UserAPIHandler.GetUserData)
	users.PUT("/update")

	jurusan := server.Group("/jurusan")
	jurusan.GET("/getAll", middleware.Auth, apiHandler.JurusanAPIHandler.GetAllJurusan)

	return r
}
