package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type Tutorial struct {
	ID		int
	Title	string 
	Author Author
	Comments []Comment
}

type Author struct {
	Name string
	Tutorials []int
}

type Comment struct {
	Body string
}

func populate() []Tutorial {
	author := &Author{Name: "Dan", Tutorials:}
	tutorial := Tutorial{
		ID: 1,
		Title: "Go Graphql Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "First Comment"},
		},
	}

	var tutorials []Tutorial
}

func main() {

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new GraphQL Schema, err %v", err)
	}

	query := `
	{
		hello
	}
		`

		params := graphql.Params{Schema: schema, RequestString: query}
		r := graphql.Do(params)
		if len(r.Errors) > 0 {
			log.Fatalf("Failed to execute graphql operation, errors: %+v", r.Errors)
		}

		rJSON, _ := json.Marshal(r)
		fmt.Printf("%s \n", rJSON)
}