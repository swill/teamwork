package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetTimeEntries() {
	// get all time_entries
	time_entries_ops := &teamwork.GetTimeEntriesOps{
		Page: 1, // required
	}
	time_entries, pages, err := conn.GetTimeEntries(time_entries_ops)
	if err != nil {
		fmt.Printf("Error getting Time Entries: %s", err.Error())
	}

	fmt.Println("\n1. Time for Company Name:", time_entries[0].CompanyName)
	fmt.Println("1. Time for Project Name:", time_entries[0].ProjectName)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectTimeEntries() {
	// get all project project_time_entries
	project_time_entries_ops := &teamwork.GetTimeEntriesOps{
		Page: 1,
	}
	project_time_entries, pages, err := conn.GetProjectTimeEntries("158721", project_time_entries_ops)
	if err != nil {
		fmt.Printf("Error getting Project Time Entries: %s", err.Error())
	}

	fmt.Println("\n1. Time for Project Name:", project_time_entries[0].ProjectName)
	fmt.Println("1. Time for Project Description:", project_time_entries[0].Description)
	fmt.Println("1. Time for Date:", project_time_entries[0].Date)
	fmt.Println("1. Time in Hours:", project_time_entries[0].Hours)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}
