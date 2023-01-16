package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/config"
	"github.com/rafli-lutfi/perpustakaan/routes"
	"github.com/rafli-lutfi/perpustakaan/utils"
)

func init() {
	config.LoadEnv()
	utils.ConnectToDatabase()
}

func main() {
	db := utils.GetDBConnection()
	r := gin.Default()

	r = routes.RunServer(db, r)

	r.Run()
}
