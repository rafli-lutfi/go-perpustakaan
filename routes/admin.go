package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/middleware"
)

func adminRoutes(router *gin.RouterGroup, handler APIHandler) {
	admin := router.Group("/admin")

	buku := admin.Group("/buku")
	buku.GET("/", middleware.Auth, handler.BukuAPIHandler.GetBukuByID)
	buku.GET("/page", middleware.Auth, handler.BukuAPIHandler.GetAllBuku)
	buku.POST("/create", middleware.Auth, handler.BukuAPIHandler.CreateBuku)
	buku.PUT("/update", middleware.Auth, handler.BukuAPIHandler.UpdateBuku)
	buku.DELETE("/delete", middleware.Auth, handler.BukuAPIHandler.DeleteBuku)

	kategori := admin.Group("/kategori")
	kategori.GET("/", middleware.Auth, handler.KategoriHandler.GetKategoriByID)
	kategori.GET("/page", middleware.Auth, handler.KategoriHandler.GetAllKategori)
	kategori.POST("/create", middleware.Auth, handler.KategoriHandler.CreateNewKategori)
	kategori.PUT("/update", middleware.Auth, handler.KategoriHandler.UpdateKategori)
	kategori.DELETE("/delete", middleware.Auth, handler.KategoriHandler.DeleteKategori)

	penerbit := admin.Group("/penerbit")
	penerbit.GET("/", middleware.Auth, handler.PenerbitHandler.GetPenerbitByID)
	penerbit.GET("/page", middleware.Auth, handler.PenerbitHandler.GetAllPenerbit)
	penerbit.POST("/create", middleware.Auth, handler.PenerbitHandler.CreateNewPenerbit)
	penerbit.PUT("/update", middleware.Auth, handler.PenerbitHandler.UpdatePenerbit)
	penerbit.DELETE("/delete", middleware.Auth, handler.PenerbitHandler.DeletePenerbit)

	author := admin.Group("/author")
	author.GET("/", middleware.Auth, handler.AuthorHandler.GetAuthorByID)
	author.GET("/page", middleware.Auth, handler.AuthorHandler.GetAllAuthor)
	author.POST("/create", middleware.Auth, handler.AuthorHandler.CreateNewAuthor)
	author.PUT("/update", middleware.Auth, handler.AuthorHandler.UpdateAuthor)
	author.DELETE("/delete", middleware.Auth, handler.AuthorHandler.DeleteAuthor)
}
