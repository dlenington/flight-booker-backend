package main 

import (
	"github.com/graphql-go/graphql"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github/com/graphql-go/graphql"
)

type Flight struct {
	ID int
	Title string
	Destination Location
	Origin Location
	Departure int
	Passengers []Passenger
}

type Location {
	ID int
	Title string
}

type Passenger {
	ID int
	FirstName string 
	LastName string
}

var flights []Flight

func populate() []Flight {
	destination := &Location{title: "San Francisco"}
	origin := &Location{title: "Minneapolis"}
	flight := Flight{
		ID: 1,
		Title: "MSP to SFO",
		Destination: *destination,
		Origin: *origin,
		Passengers: []Passenger{
			Passenger{ID: 1, FirstName: "Ben", LastName: "L"},
			Passenger{ID: 2, FirstName: "Dan", LastName: "L"}
		}
	}
	flight2 := Flight{
		ID: 2,
		Title: "MSP to LAX",
		Destination: *destination,
		Origin: *origin,
		Passengers: []Passenger{
			Passenger{ID: 1, FirstName: "Ben", LastName: "L"},
			Passenger{ID: 2, FirstName: "Dan", LastName: "L"}
			Passenger{ID: 3, FirstName: "Mom", LastName: "L"},
			Passenger{ID: 4, FirstName: "Dad", LastName: "L"},
		}
	}
	var flights []Flight
	flights = append(flights, flight)
	flights = append(flights, flight2)

	return flights
}

//Defining types

var locationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int
			},
			"title": &graphql.Field{
				Type: graphql.String
			}
		}
	}
)

var passengerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int
			},
			"firstName": &graphql.Field{
				Type: graphql.String
			},
			"lastName": &graphql.Field{
				Type: graphql.String
			}
		}
	}
)

var flightType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Flight",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int
			},
			"title": &graphql.Field{
				Type: graphql.String
			},
			"destination": &graphql.Field{
				Type: locationType,
			},
			"origin": &graphql.Field{
				Type: locationType,
			},
			"passengers": &graphql.Field{
				Type: graphql.NewList(passengerType),
			}
		}
	}
)

func main() {
	tutorials = populate()

	fields := graphql.Field{
		"flight": &graphql.Field{
			Type: tutorialType,
			Description: "Get Flight By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println("Flight queried")
			}
		},

		"list": &graphql.Field{
			Type: graphql.NewList(tutorialType),
			Description: "Get Flight List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var list []Flight
				return list
			}
		},

	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		Mutation: mutationType,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create newd GraphQL Schema, err %v", err)
	}

	query := `
	mutation {
		create(title: "Hello World")  {
			
			
			}`
}