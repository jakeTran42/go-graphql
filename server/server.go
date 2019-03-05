package server

import (
	"encoding/json"
	"net/http"

	"github.com/jakeTran42/go-graphql/graphql"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

// Server hold connection to the db
type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

// GraphQL returns an http.HandlerFunc for for /graphql
func (s *Server) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//edge case
		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)

		// render.JSON comes from the chi/render package
		render.JSON(w, r, result)
	}
}
