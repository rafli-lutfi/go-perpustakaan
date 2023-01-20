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
	BukuAPIHandler    controllers.BukuAPI
	KategoriHandler   controllers.KategoriAPI
	PenerbitHandler   controllers.PenerbitAPI
	AuthorHandler     controllers.AuthorAPI
}

func RunServer(db *gorm.DB, r *gin.Engine) *gin.Engine {
	userDatabase := database.NewUserDatabase(db)
	jurusanDatabase := database.NewJurusanDatabase(db)
	bukuDatabase := database.NewBukuDatabase(db)
	kategoriDatabase := database.NewKategoriDatabase(db)
	penerbitDatabase := database.NewPenerbitDatabase(db)
	authorDatabase := database.NewAuthorDatabase(db)

	userService := service.NewUserService(userDatabase, jurusanDatabase)
	jurusanService := service.NewJurusanService(jurusanDatabase)
	bukuService := service.NewBukuService(bukuDatabase, kategoriDatabase, penerbitDatabase, authorDatabase)
	kategoriService := service.NewKategoriService(kategoriDatabase)
	penerbitService := service.NewPenerbitService(penerbitDatabase)
	authorService := service.NewAuthorService(authorDatabase)

	userAPIHandler := controllers.NewUserAPI(userService)
	jurusanAPIHandler := controllers.NewJurusanAPI(jurusanService)
	bukuAPIHandler := controllers.NewBukuAPI(bukuService)
	kategoriAPIHandler := controllers.NewKategoriAPI(kategoriService)
	penerbitAPIHandler := controllers.NewPenerbitAPI(penerbitService)
	authorAPIHandler := controllers.NewAuthorAPI(authorService)

	apiHandler := APIHandler{
		UserAPIHandler:    userAPIHandler,
		JurusanAPIHandler: jurusanAPIHandler,
		BukuAPIHandler:    bukuAPIHandler,
		KategoriHandler:   kategoriAPIHandler,
		PenerbitHandler:   penerbitAPIHandler,
		AuthorHandler:     authorAPIHandler,
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
	jurusan.POST("/create", middleware.Auth, apiHandler.JurusanAPIHandler.CreateJurusan)
	jurusan.PUT("/update", middleware.Auth, apiHandler.JurusanAPIHandler.UpdateJurusan)
	jurusan.DELETE("/delete", middleware.Auth, apiHandler.JurusanAPIHandler.DeleteJurusan)

	adminRoutes(server, apiHandler)

	return r
}
