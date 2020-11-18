package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 微信接口配置验证
func Index(c *gin.Context) {
	c.String(http.StatusOK, "123")
}
