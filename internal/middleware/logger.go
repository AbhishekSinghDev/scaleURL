package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// runs pre-request
		start := time.Now()

		ctx.Next()

		// runs post-request
		log.Info().Str("method", ctx.Request.Method).Str("path", ctx.Request.URL.Path).Int("status", ctx.Writer.Status()).Dur("latency", time.Since(start)).Msg("request completed")
	}
}
