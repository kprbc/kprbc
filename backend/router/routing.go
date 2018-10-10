package router

import "github.com/gin-gonic/gin"

// Route adds routes to a *gin.Engine server
func Route(s *gin.Engine) *gin.Engine {
	s.GET("/healthz", healthCheckHandler)
	return s
}
