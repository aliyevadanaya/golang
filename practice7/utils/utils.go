package utils

import (
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "no token"})
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if token == nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("userID", claims["user_id"].(string))
		c.Set("role", claims["role"].(string))

		c.Next()
	}
}

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, _ := c.Get("role")

		if r != role {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}

type clientData struct {
	count     int
	timestamp int64
}

var requests = make(map[string]clientData)
var mutex sync.Mutex

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()

		if id, ok := c.Get("userID"); ok {
			key = id.(string)
		}

		now := time.Now().Unix()

		mutex.Lock()
		data := requests[key]

		if now-data.timestamp > 30 {
			data.count = 0
			data.timestamp = now
		}

		data.count++
		requests[key] = data
		count := data.count
		mutex.Unlock()

		if count > 10 {
			c.AbortWithStatusJSON(429, gin.H{"error": "too many requests"})
			return
		}

		c.Next()
	}
}
