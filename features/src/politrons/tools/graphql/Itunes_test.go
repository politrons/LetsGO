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
This funcrio
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
			"songs": loadSongsField(),
			"album": loadAlbumField(),
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
	Args:The that we expect to receive in the query
	Resolve: Function to be invoked when the query use this field
	Description: description of the field

*/
func loadSongsField() *graphql.Field {
	return &graphql.Field{
		Name: "Songs filed",
		Type: graphql.NewList(songType),
		Args: graphql.FieldConfigArgument{
			"album": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve:     songResolver,
		Description: "Songs field to find a particular song using album",
	}
}

func loadAlbumField() *graphql.Field {
	return &graphql.Field{
		Type: albumType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: albumResolver,
	}
}

//######################
//	GRAPHQL RESOLVERS  #
//######################

var songResolver = func(params graphql.ResolveParams) (interface{}, error) {
	album := params.Args["album"].(string)
	println(album)
	return songs, nil
}

var albumResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(string)
	for _, album := range albums {
		if album.ID == id {
			return album, nil
		}
	}
	return nil, nil
}

//#################
//	  GO TYPES    #
//#################
type Album struct {
	ID     string `json:"id,omitempty"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	Type   string `json:"type"`
}

type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

//#########################
//	GRAPHQL OBJECT TYPES  #
//#########################

var songType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Song",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"album": &graphql.Field{
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

var artistType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Artist",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var albumType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Album",
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

var albums []Album = []Album{
	Album{
		ID:     "ts-fearless",
		Artist: "1",
		Title:  "Fearless",
		Year:   "2008",
		Type:   "album",
	},
}

var artists []Artist = []Artist{
	Artist{
		ID:   "1",
		Name: "Taylor Swift",
		Type: "artist",
	},
}

var songs []Song = []Song{
	Song{
		ID:       "1",
		Album:    "ts-fearless",
		Title:    "Fearless",
		Duration: "4:01",
		Type:     "song",
	},
	Song{
		ID:       "2",
		Album:    "ts-fearless",
		Title:    "Fifteen",
		Duration: "4:54",
		Type:     "song",
	},
}

/*curl -g 'http://localhost:12345/graphql?query={songs(album:"ts-fearless"){title,duration}}'
 */
