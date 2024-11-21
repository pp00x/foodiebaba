package middlewares

import (
    "net/http"
    "os"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            return
        }
        tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("userID", uint(claims["user_id"].(float64)))
            c.Set("userRole", claims["role"].(string))
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }
        c.Next()
    }
}