package server

import (
	"brendisurfs/crystalshop-backend/server/handlers"

	"github.com/gin-gonic/gin"
)

func Server() {

	r := gin.Default()

	r.GET("/api/crystals", handlers.CrystalHandler)

	// use cors for local dev.
	r.Use()

	r.Run()
}
