package main

import (
	controller "Jurusan/Controller"
	connection "Jurusan/Model/Connection"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":1213"
	r := gin.Default()
	connection.ConnectDatabase()

	//###BEGIN WEB API
	// Get data
	r.GET("/api/Jurusan/GetJurusan", controller.GetJurusanByID)

	// Get data
	r.GET("/api/Jurusan/GetJurusan/Id", controller.GetJurusanByID)

	// Insert data
	r.POST("/api/Jurusan/InsertJurusan", controller.InsertJurusan)

	// Update data
	r.PUT("/api/Jurusan/UpdateJurusan", controller.UpdateJurusan)

	// Delete data
	r.DELETE("/api/Jurusan/DeleteJurusan", controller.DeleteJurusan)
	//###END WEB API

	r.Run(port)
}
