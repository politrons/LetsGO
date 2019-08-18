package graphql

import (
	"log"
	"net/http"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

/**
A GraphQL query is performed against this schema with the resulting output printed in JSON format
*/
func TestRunMainServer(t *testing.T) {
	s := `
                schema {
                        query: Query
						subscription: Subscription
                }
                type Query {
                        user: String!
                }
				type Subscription {
					helloSaid(): Foo!
				}
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type query struct {
	id string
}

type subscription struct {
	id string
}

type Foo struct {
	id string
}

func (q *query) HelloSaid() Foo {
	return Foo{"Hello, user!"}
}

func (q *query) User() string {
	return "Hello, user!"
}
