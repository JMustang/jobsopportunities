package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default() // <-Inicializa a ROTA

	// As ROTAS seram inicializadas aqui
	initializeRoutes(r)

	r.Run(":8080") // Roda o servifor
}
