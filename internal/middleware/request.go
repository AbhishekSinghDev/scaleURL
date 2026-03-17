package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := uuid.New().String()
		// sets in gin context
		ctx.Set("request_id", id)
		// sets in header
		ctx.Writer.Header().Set("SURL-REQUEST-ID", id)
		ctx.Next()
	}
}
