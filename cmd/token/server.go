package token

import (
	"authentication-service/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewServer(config *config.AppConfigs, logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "Token REST API Server",
		Long:  "Token management service server",
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()
			r.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
			r.Run(fmt.Sprintf(":%d", config.TokenServerConfig.Port))
		},
	}

	return cmd
}
