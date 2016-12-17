package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetPeople() {
	// get all people
	True := true
	people_ops := &teamwork.GetPeopleOps{
		FullProfile: &True,
	}
	people, pages, err := conn.GetPeople(people_ops)
	if err != nil {
		fmt.Printf("Error getting People: %s", err.Error())
	}

	fmt.Println("GetPeople")
	fmt.Println("1. Username:", people[0].UserName)
	fmt.Println("1. Full Name:", people[0].FirstName, people[0].LastName)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectPeople() {
	// get project people
	people_ops = &teamwork.GetPeopleOps{
		FullProfile: &True,
	}
	people, pages, err = conn.GetProjectPeople("158721", people_ops)
	if err != nil {
		fmt.Printf("Error getting Project People: %s", err.Error())
	}

	fmt.Println("\nGetProjectPeople")
	fmt.Println("1. Username:", people[0].UserName)
	fmt.Println("1. Full Name:", people[0].FirstName, people[0].LastName)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetCompanyPeople() {
	// get company people
	people_ops = &teamwork.GetPeopleOps{
		FullProfile: &True,
	}
	people, pages, err = conn.GetCompanyPeople(conn.Account.CompanyID, people_ops)
	if err != nil {
		fmt.Printf("Error getting Company People: %s", err.Error())
	}

	fmt.Println("\nGetCompanyPeople")
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

	fmt.Println("GetPerson")
	fmt.Println("Username:", person.UserName)
	fmt.Println("Full Name:", person.FirstName, person.LastName)
}
