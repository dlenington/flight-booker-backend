package main 

import (
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