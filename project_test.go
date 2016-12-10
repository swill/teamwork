package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetProjects() {
	// get all projects
	projects_ops := &teamwork.GetProjectsOps{
		Status: "ALL",
	}
	projects, pages, err := conn.GetProjects(projects_ops)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	fmt.Println("1. ID: ", projects[0].ID)
	fmt.Println("1. Name: ", projects[0].Name)
	fmt.Println("1. Status: ", projects[0].Status)
	fmt.Println("on page #: ", pages.Page)
	fmt.Println("# of pages: ", pages.Pages)
	fmt.Println("# of records: ", pages.Records)
}

func ExampleConnection_GetProject() {
	// get one project
	True := true
	project_ops := &teamwork.GetProjectOps{
		IncludePeople: &True,
	}
	project, err := conn.GetProject("158747", project_ops)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	fmt.Println("ID:", project.ID)
	fmt.Println("Name:", project.Name)
	fmt.Println("Status:", project.Status)
}
