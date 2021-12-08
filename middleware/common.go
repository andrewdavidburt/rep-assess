package middleware

import ("github.com/gin-gonic/gin"
		"github.com/google/uuid"
)

// GenerateRequestID is a piece of middleware that attaches a request id to the request for traceability in logs
func GenerateRequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		ID := uuid.New()
		context.Header("X-Request-ID", ID.String())
		context.Next()
	}
}