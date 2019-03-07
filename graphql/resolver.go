package graphql

import (
	"github.com/jakeTran42/go-graphql/postgres"
	"github.com/graphql-go/graphql"
)

// Resolver struct connect to database
type Resolver struct {
	db *postgres.Db
}

// UserRes call GetUser to query users by name
func (r *Resolver) UserRes(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	name, ok := p.Args["name"].(string)
	if ok {
		users := r.db.GetUser(name)
		return users, nil
	}

	return nil, nil
}