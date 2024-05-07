package http

import (
	"context"
	"net/http"
	"time"

	models "github.com/AbnerCMoura/WEB_API_GOLANG/internal/Models"
	repositories "github.com/AbnerCMoura/WEB_API_GOLANG/internal/Repositories"
	services "github.com/AbnerCMoura/WEB_API_GOLANG/internal/Services"
	"github.com/AbnerCMoura/WEB_API_GOLANG/internal/database"
	"github.com/gin-gonic/gin"
)

var service services.UsuarioService

func Configure() {
	service = services.UsuarioService{
		UsuarioRepository: &repositories.RepositoryDb{
			Connection: database.Connection,
		},
	}
}

func AdicionarUsuiario(ctx *gin.Context) {
	var usuario models.Usuario

	if err := ctx.BindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	response, err := service.Inserir(ctxTimeout, usuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
