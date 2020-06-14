package http_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func systemHealth(c *gin.Context, _ *GinServer) {
	c.JSON(http.StatusOK, &GinResponse{Data: "OK"})
}
