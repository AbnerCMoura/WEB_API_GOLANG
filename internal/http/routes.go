package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.POST("/adicioar", AdicionarUsuiario)
	// g.DELETE("/deletar/:id", DeletePosts)
	// g.GET("/pegar/:id", GetPosts)
	// g.GET("/pegartodos", GetAll)
	// g.PUT("/atualizar/:id", Update)

	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
