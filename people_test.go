package teamwork_test

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetPeople() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all people
	True := true
	peopleOps := &teamwork.GetPeopleOps{
		FullProfile: &True,
	}
	people, pages, err := conn.GetPeople(peopleOps)
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
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get project people
	True := true
	peopleOps := &teamwork.GetPeopleOps{
		FullProfile: &True,
	}
	people, pages, err := conn.GetProjectPeople("158721", peopleOps)
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
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get company people
	True := true
	peopleOps := &teamwork.GetPeopleOps{
		FullProfile: &True,
	}
	people, pages, err := conn.GetCompanyPeople(conn.Account.CompanyID, peopleOps)
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
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get one person
	person, err := conn.GetPerson("85457")
	if err != nil {
		fmt.Printf("Error getting Person: %s", err.Error())
	}

	fmt.Println("GetPerson")
	fmt.Println("Username:", person.UserName)
	fmt.Println("Full Name:", person.FirstName, person.LastName)
}
