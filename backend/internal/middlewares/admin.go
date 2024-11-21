package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("userRole")
        if !exists || role != "admin" {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admins only"})
            return
        }
        c.Next()
    }
}