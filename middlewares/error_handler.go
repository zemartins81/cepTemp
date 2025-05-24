package middlewares

import (
	"net/http"

	"cepTemp/models"
	"github.com/gin-gonic/gin"
)

// ErrorHandler é um middleware para tratamento centralizado de erros
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Verifica se houve algum erro durante o processamento da requisição
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Trata diferentes tipos de erro
			switch err.Error() {
			case "invalid zipcode":
				c.JSON(http.StatusUnprocessableEntity, models.ErrorResponse{
					Message: "invalid zipcode",
				})
			case "can not find zipcode":
				c.JSON(http.StatusNotFound, models.ErrorResponse{
					Message: "can not find zipcode",
				})
			default:
				c.JSON(http.StatusInternalServerError, models.ErrorResponse{
					Message: "internal server error",
				})
			}

			// Interrompe o processamento
			c.Abort()
		}
	}
}
