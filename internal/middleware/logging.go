package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logging(c *gin.Context, logger *zap.Logger) {
	logger.Info("request received",
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
	)

	c.Next()

	logger.Info("response sent",
		zap.Int("status", c.Writer.Status()),
	)
}
