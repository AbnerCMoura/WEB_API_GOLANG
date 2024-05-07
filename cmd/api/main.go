package main

import (
	"context"

	"github.com/AbnerCMoura/WEB_API_GOLANG/internal/database"
	"github.com/AbnerCMoura/WEB_API_GOLANG/internal/http"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	connectionString := "postgresql://localhost:5432/"

	connectionDb, errConnection := database.NewConnection(ctx, connectionString)
	if errConnection != nil {
		panic(errConnection)
	}

	defer connectionDb.Close()

	g := gin.New()
	g.Use(gin.Recovery())
	http.Configure()
	http.SetRoutes(g)
	g.Run(":8090")
}
