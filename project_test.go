package teamwork_test

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetProjects() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all projects
	projectsOps := &teamwork.GetProjectsOps{
		Status: "ALL",
	}
	projects, pages, err := conn.GetProjects(projectsOps)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	fmt.Println("GetProjects")
	fmt.Println("1. Name:", projects[0].Name)
	fmt.Println("1. Status:", projects[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProject() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get one project
	True := true
	projectOps := &teamwork.GetProjectOps{
		IncludePeople: &True,
	}
	project, err := conn.GetProject("158747", projectOps)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	fmt.Println("GetProject")
	fmt.Println("ID:", project.ID)
	fmt.Println("Name:", project.Name)
	fmt.Println("Status:", project.Status)
}
