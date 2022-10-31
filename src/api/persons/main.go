package main

import (
	"bank/api/persons/api/controllers"
	"bank/api/persons/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Migrate()
	startServer()
}

func startServer() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		controllers.Routes(v1)
	}
	var err = router.Run()

	if err != nil {
		panic(err)
	}
}
