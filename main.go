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
	flight := Flight{
		ID: 1,
		Title: "MSP to SFO",
		Destination: *destination
	}
}