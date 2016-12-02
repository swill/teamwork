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

	params := &teamwork.ProjectsOps{
		Status: "ALL",
	}
	projects, err := conn.GetProjects(params)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	fmt.Println("1. Name: ", projects[0].Name)
	fmt.Println("1. Status: ", projects[0].Status)
}
```