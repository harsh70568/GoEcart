package main

import (
	"goEcart/db"
	"goEcart/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	/* Connect to the DB */
	db.ConnectDB()

	/* Set up router */
	router := gin.Default()

	/* Set up routing */
	routes.AdminRoutes(router)
	routes.UserRoutes(router)

	/* Run the server */
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8080"
	}
	router.Run(PORT)
}
