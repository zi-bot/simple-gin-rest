package middleware

import "github.com/gin-gonic/gin"

func Recover() gin.HandlerFunc {
	return gin.Recovery()
}
