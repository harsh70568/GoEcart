package main

import (
	"fmt"
	"goEcart/db"
	"goEcart/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading the env file...")
	}
	PORT := viper.GetString("PORT")
	if PORT == "" {
		PORT = ":8080"
	}
	router.Run(PORT)
}
