package main

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/graphql-go/graphql"
)



type Flight struct {
	ID          string
	Title       string
	Destination Location
	Origin      Location
	Departure   int
	Passengers  []Passenger
}

type Location struct {
	ID    string
	Title string
}

type Passenger struct {
	ID        string
	FirstName string
	LastName  string
}

var flights []Flight

func populate() []Flight {
	destination := &Location{Title: "San Francisco"}
	origin := &Location{Title: "Minneapolis"}
	flight := Flight{
		ID:          "1",
		Title:       "MSP to SFO",
		Destination: *destination,
		Origin:      *origin,
		Passengers: []Passenger{
			Passenger{ID: "1", FirstName: "Ben", LastName: "L"},
			Passenger{ID: "2", FirstName: "Dan", LastName: "L"},
		},
	}
	flight2 := Flight{
		ID:          "2",
		Title:       "MSP to LAX",
		Destination: *destination,
		Origin:      *origin,
		Passengers: []Passenger{
			Passenger{ID: "1", FirstName: "Ben", LastName: "L"},
			Passenger{ID: "2", FirstName: "Dan", LastName: "L"},
			Passenger{ID: "3", FirstName: "Mom", LastName: "L"},
			Passenger{ID: "4", FirstName: "Dad", LastName: "L"},
		},
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
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var passengerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Passenger",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var flightType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Flight",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"destination": &graphql.Field{
				Type: locationType,
			},
			"origin": &graphql.Field{
				Type: locationType,
			},
			"passengers": &graphql.Field{
				Type: graphql.NewList(passengerType),
			},
		},
	},
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
        fmt.Println(err.Error())
        return
	}
	
	svc := dynamodb.New(sess)
	
	


	flights = populate()

	fields := graphql.Fields{
		"flight": &graphql.Field{
			Type:        flightType,
			Description: "Get Flight By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				tableName := "rockmed-api-SampleTable-A8FNI2HZFC56"
				id := p.Args["id"].(string)
				result, err := svc.GetItem(&dynamodb.GetItemInput{
					TableName: aws.String(tableName),
					Key: map[string]*dynamodb.AttributeValue{
						"id": {
							S: aws.String(id),
						},
					},
				})

				if err != nil {
					fmt.Printf(err.Error())
					return nil, nil
				} 

				if result.Item == nil {
					msg := "Could not find "
					return nil, errors.New(msg)
				}

				flight := Flight{}

				err = dynamodbattribute.UnmarshalMap(result.Item, &flight)
				if err != nil {
					panic(fmt.Sprintf("Failed to unmarshall Record, %v, err "))
				}

				fmt.Println("Found flight:")
				fmt.Println("Name: ", flight.Title)
				fmt.Println("Flight queried")
				return nil, nil
			},
		},

		"list": &graphql.Field{
			Type:        graphql.NewList(flightType),
			Description: "Get Flight List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				//Use dynamodb.Query method
			tableName := "rockmed-api-SampleTable-A8FNI2HZFC56"
			input := &dynamodb.QueryInput{
				TableName := aws.String(tableName)
			}
			result, err := svc.Query(input)
			if err != nil {
				fmt.Printf("ERROR: %v\n", err.Error())
				return
			}

			fmt.Println(result)
				
				var list []Flight
				return list, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new GraphQL Schema, err %v", err)
	}

	query := `
	query {
	flight(id:"1234ABCD") {
		title
	}
}
			`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	//Query
}
