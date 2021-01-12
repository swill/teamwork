package teamwork_test

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetTimeEntries() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all time entries
	Page := 200
	timeEntriesOps := &teamwork.GetTimeEntriesOps{
		Page: &Page,
	}
	timeEntries, pages, err := conn.GetTimeEntries(timeEntriesOps)
	if err != nil {
		fmt.Printf("Error getting Time Entries: %s", err.Error())
	}

	fmt.Println("GetTimeEntries")
	fmt.Println("1. Time for Company Name:", timeEntries[0].CompanyName)
	fmt.Println("1. Time for Project Name:", timeEntries[0].ProjectName)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectTimeEntries() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all project time entries
	projectTimeEntriesOps := &teamwork.GetTimeEntriesOps{}
	projectTimeEntries, pages, err := conn.GetProjectTimeEntries("158721", projectTimeEntriesOps)
	if err != nil {
		fmt.Printf("Error getting Project Time Entries: %s", err.Error())
	}

	fmt.Println("GetProjectTimeEntries")
	fmt.Println("1. Time for Project Name:", projectTimeEntries[0].ProjectName)
	fmt.Println("1. Time for Project Description:", projectTimeEntries[0].Description)
	fmt.Println("1. Time for Date:", projectTimeEntries[0].Date)
	fmt.Println("1. Time in Hours:", projectTimeEntries[0].Hours)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
}

func ExampleConnection_GetTaskTimeEntries() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all task time entries
	taskTimeEntriesOps := &teamwork.GetTimeEntriesOps{}
	taskTimeEntries, pages, err := conn.GetTaskTimeEntries("4754100", taskTimeEntriesOps)
	if err != nil {
		fmt.Printf("Error getting Task Time Entries: %s", err.Error())
	}

	fmt.Println("GetTaskTimeEntries")
	fmt.Println("1. Time for Task List:", taskTimeEntries[0].TaskListName)
	fmt.Println("1. Time for Task Name:", taskTimeEntries[0].TaskItemName)
	fmt.Println("1. Time for Date:", taskTimeEntries[0].Date)
	fmt.Println("1. Time in Hours:", taskTimeEntries[0].Hours)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetTotalTime() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get total time for an account
	totalTimeOps := &teamwork.GetTotalTimeOps{}
	totalTime, err := conn.GetTotalTime(totalTimeOps)
	if err != nil {
		fmt.Printf("Error getting Total Time: %s", err.Error())
	}

	fmt.Println("GetTotalTime")
	fmt.Println("Total Hours:", totalTime.TotalHoursSum)
	fmt.Println("Total Hours Billable:", totalTime.BillableHoursSum)
}

func ExampleConnection_GetProjectTotalTime() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get total time for a project
	totalTimeOps := &teamwork.GetTotalTimeOps{}
	projectTotalTime, err := conn.GetProjectTotalTime("158721", totalTimeOps)
	if err != nil {
		fmt.Printf("Error getting Project Total Time: %s", err.Error())
	}

	fmt.Println("GetProjectTotalTime")
	fmt.Println("Name:", projectTotalTime[0].Name)
	fmt.Println("Total Hours:", projectTotalTime[0].TimeTotals.TotalHoursSum)
	fmt.Println("Total Hours Billable:", projectTotalTime[0].TimeTotals.BillableHoursSum)
}

func ExampleConnection_GetTaskListTotalTime() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get total time for a task list
	totalTimeOps := &teamwork.GetTotalTimeOps{}
	taskListTotalTime, err := conn.GetTaskListTotalTime("704748", totalTimeOps)
	if err != nil {
		fmt.Printf("Error getting Project Total Time: %s", err.Error())
	}

	fmt.Println("GetTaskListTotalTime")
	fmt.Println("Name:", taskListTotalTime[0].TaskList.Name)
	fmt.Println("Total Hours:", taskListTotalTime[0].TaskList.TimeTotals.TotalHoursSum)
	fmt.Println("Total Hours Billable:", taskListTotalTime[0].TaskList.TimeTotals.BillableHoursSum)
}

func ExampleConnection_GetTaskTotalTime() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get total time for a task
	totalTimeOps := &teamwork.GetTotalTimeOps{}
	taskTotalTime, err := conn.GetTaskTotalTime("4486838", totalTimeOps)
	if err != nil {
		fmt.Printf("Error getting Project Total Time: %s", err.Error())
	}

	fmt.Println("GetTaskTotalTime")
	fmt.Println("Name:", taskTotalTime[0].TaskList.Task.Name)
	fmt.Println("Total Hours:", taskTotalTime[0].TaskList.Task.TimeTotals.TotalHoursSum)
	fmt.Println("Total Hours Billable:", taskTotalTime[0].TaskList.Task.TimeTotals.BillableHoursSum)
}
