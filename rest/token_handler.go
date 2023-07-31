package rest

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken(ctx *gin.Context) (bool, error) {

}

func getTokenFromRedis(c *gin.Context) (*string, error) {
	return nil, nil
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
