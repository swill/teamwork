TeamWork
========
[![GoDoc](https://godoc.org/github.com/swill/teamwork?status.svg)](https://godoc.org/github.com/swill/teamwork)

It is still VERY early for this project, so please keep that in mind.  I am working out how I want to manage all of the parameters and configuration options in a generalized way and then I will start adding more components.

Feedback is welcome even at this very early stage.

Here is an example application using this library.
```go
package main

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

var (
	conn      *teamwork.Connection
	api_token = "my api token"
)

func main() {
	// setup the teamwork connection
	conn, err := teamwork.Connect(api_token)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all projects
	projects_ops := &teamwork.ProjectsOps{
		Status: "ALL",
	}
	projects, pages, err := conn.GetProjects(projects_ops)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	fmt.Println("1. Name: ", projects[0].Name)
	fmt.Println("1. Status: ", projects[0].Status)
	fmt.Println("on page #: ", pages.Page)
	fmt.Println("# of pages: ", pages.Pages)
	fmt.Println("# of records: ", pages.Records)

	// get one project
	project_ops := &teamwork.ProjectOps{}
	project, err := conn.GetProject("158747", project_ops)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	fmt.Println("\nName: ", project.Name)
	fmt.Println("Status: ", project.Status)
}
```