package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetPeople() {
	// get all people
	people_ops := &teamwork.GetPeopleOps{
		FullProfile:      true,
		ReturnProjectIds: true,
	}
	people, pages, err := conn.GetPeople(people_ops)
	if err != nil {
		fmt.Printf("Error getting People: %s", err.Error())
	}

	fmt.Println("1. Username:", people[0].UserName)
	fmt.Println("1. Full Name:", people[0].FirstName, people[0].LastName)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetPerson() {
	// get one person
	person, err := conn.GetPerson("85457")
	if err != nil {
		fmt.Printf("Error getting Person: %s", err.Error())
	}

	fmt.Println("Username:", person.UserName)
	fmt.Println("Full Name:", person.FirstName, person.LastName)
}
