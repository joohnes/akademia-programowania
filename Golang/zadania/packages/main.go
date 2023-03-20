package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/grupawp/appdispatcher"
)

type Student struct {
	FirstName     string
	LastName      string
	applicationID uuid.UUID
}

func (s Student) ApplicationID() string {
	return s.applicationID.String()
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func main() {
	student := Student{"Kuba", "Pienkowski", uuid.New()}
	status, err := appdispatcher.Submit(student)
	if err != nil {
		log.Println(err)
	}
	log.Println(status, err)
}
