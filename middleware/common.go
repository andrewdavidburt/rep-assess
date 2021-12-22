package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "reperio-backend-assessment/database"
	// "reperio-backend-assessment/models"
	// "reperio-backend-assessment/packages"
)

// GenerateRequestID is a piece of middleware that attaches a request id to the request for traceability in logs
func GenerateRequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		ID := uuid.New()
		context.Header("X-Request-ID", ID.String())
		context.Next()
	}
}

// IsCurrentLocationPresent is middleware that checks whether the db has the request already present
func IsCurrentLocationPresent() gin.HandlerFunc {
	return func(context *gin.Context) {
		// var CLPresent bool
		// fmt.Println("IsCurrentLocationPresent")
		// context.Keys["CLPresent"] = CLPresent
		context.Next()
	}
}
