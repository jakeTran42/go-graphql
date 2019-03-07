package graphql

import (
	"github.com/jakeTran42/go-graphql/postgres"
	"github.com/graphql-go/graphql"
)

// Root struct for querying
type Root struct {
	Query *graphql.Object
}

// NewRoot returns root query type
func NewRoot(db *postgres.Db) *Root {
	// Our resolvers to db receiving data
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						// Slice of User type
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserRes,
					},
				},
			},
		),
	}
	return &root
}