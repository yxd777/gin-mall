package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	c.String(http.StatusOK, "hello Profile")
}
