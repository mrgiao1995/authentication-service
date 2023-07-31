package auth

import (
	"authentication-service/config"
	"authentication-service/graph"
	"authentication-service/graph/resolver"
	"authentication-service/model"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func NewServer(config *config.AppConfigs, logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Auth GraphQL Server",
		Long:  "Authentiacation service GraphQL server",
		Run: func(cmd *cobra.Command, args []string) {
			dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
				config.DbConfig.Host,
				config.DbConfig.UserName,
				config.DbConfig.Password,
				config.DbConfig.Name,
				config.DbConfig.Port)
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   fmt.Sprintf("%s.", config.DbConfig.Schema), // schema name
					SingularTable: false,
				}})
			if err != nil {
				log.Fatal(err)
			}

			db.AutoMigrate(&model.User{})

			r := resolver.NewResolver(config, logger, db)
			c := graph.Config{Resolvers: r}
			srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

			http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
			http.Handle("/graphql", srv)

			log.Printf("connect to http://%s:%d for GraphQL playground", config.AuthServerConfig.Host, config.AuthServerConfig.Port)
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.AuthServerConfig.Port), nil))
		},
	}

	return cmd
}
