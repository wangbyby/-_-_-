package controller

import (
	"github.com/gin-gonic/gin"
)

//统一注册路由
func RegisterRoutes(g *gin.RouterGroup) {
	new(IndexController).RegisterRoute(g)
}
