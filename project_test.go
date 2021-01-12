package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetProjects() {
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
