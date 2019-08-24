package graphql

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestItunes(t *testing.T) {
	http.HandleFunc("/graphql", handleQuery())
	_ = http.ListenAndServe(":12345", nil)
}

/**
This
*/
func handleQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        createSchema(),
			RequestString: r.URL.Query().Get("query"),
		})
		err := json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Println("Error Serializing result")
			panic(err)
		}
	}
}

/**
THis function use factory [graphql.NewSchema] where we pass [SchemaConfig] which basically require
a Object Query type
*/
func createSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: createQuery(),
	})
	if err != nil {
		log.Println("Error creating Schema")
		panic(err)
	}
	return schema
}

/**
Graphql query object it's created using [graphql.NewObject] where we pass a [ObjectConfig]
this ObjectConfig it's created filling with:

	Name: Name of the Query
	Fields: map [string]Field where each Field is a select field in your query
	Description: string description of what the query is about.

*/
func createQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"actors": loadActorsField(),
			"movie":  loadMovieField(),
		},
		Description: "A movie query with information of most famous movies",
	})
}

//######################
//	 GRAPHQL FIELDS    #
//######################

/**
A Field in graphql is through the type [Field] which is expected to be filled with:
	Name:Name of the field
	Type: Object type that define the attributes
	Args:The filter arg that we expect to receive in the query and we will use in [Resolve] function
	Resolve: Function to be invoked when the query use this field [actors] and is where use the args to filter from all the data
	Description: description of the field

*/
func loadActorsField() *graphql.Field {
	return &graphql.Field{
		Name: "Actor filed",
		Type: graphql.NewList(actorType),
		Args: graphql.FieldConfigArgument{
			"movie": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve:     actorResolver,
		Description: "Songs field to find a particular song using album",
	}
}

func loadMovieField() *graphql.Field {
	return &graphql.Field{
		Type: movieType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: movieResolver,
	}
}

//######################
//	GRAPHQL RESOLVERS  #
//######################

var actorResolver = func(params graphql.ResolveParams) (interface{}, error) {
	movie := params.Args["movie"].(string)
	filterActors := []Actor{}
	for i := 0; i < len(actors); i++ {
		if actors[i].Movie == movie {
			filterActors = append(filterActors, actors[i])
		}
	}
	/*	for _, song := range actors {
		if song.ID == movie {
			filterActors = append(filterActors, song)
		}
	}*/
	return filterActors, nil
}

var movieResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)
	for _, movie := range albums {
		if movie.ID == id {
			return movie, nil
		}
	}
	return nil, nil
}

//#################
//	  GO TYPES    #
//#################
type Movie struct {
	ID     string `json:"id,omitempty"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	Type   string `json:"type"`
}

type Actor struct {
	ID       string `json:"id,omitempty"`
	Movie    string `json:"movie"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

//#########################
//	GRAPHQL OBJECT TYPES  #
//#########################

var actorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Actor",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"movie": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"duration": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var movieType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Movie",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"artist": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"genre": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

//##################
//	  MOCK DATA    #
//##################

var albums = []Movie{
	{
		ID:     "ts-fearless",
		Artist: "1",
		Title:  "Fearless",
		Year:   "2008",
		Type:   "album",
	},
}

var actors = []Actor{
	{
		ID:       "1",
		Movie:    "Titanic",
		Title:    "Fearless",
		Duration: "4:01",
		Type:     "song",
	},
	{
		ID:       "2",
		Movie:    "ts-fearless",
		Title:    "Fifteen",
		Duration: "4:54",
		Type:     "song",
	},
}

/*curl -g 'http://localhost:12345/graphql?query={actors(movie:"ts-fearless"){title,duration}}'
 */
