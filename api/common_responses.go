package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondWithValidationError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, nil)
}
