package teamwork_test

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetTasks() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all tasks
	True := true
	PageSize := 250
	Page := 2
	tasksOps := &teamwork.GetTasksOps{}
	tasksOps.GetFiles = &True
	tasksOps.IncludeCompletedSubtasks = &True
	tasksOps.IncludeCompletedTasks = &True
	tasksOps.NestSubTasks = "yes"
	tasksOps.PageSize = &PageSize
	tasksOps.Page = &Page
	tasks, pages, err := conn.GetTasks(tasksOps)
	if err != nil {
		fmt.Printf("Error getting Tasks: %s", err.Error())
	}

	fmt.Println("GetTasks")
	fmt.Println("1. Task List Name:", tasks[0].TaskListName)
	fmt.Println("1. Task ID:", tasks[0].ID)
	fmt.Println("1. Task Content:", tasks[0].Content)
	fmt.Println("1. Task Status:", tasks[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectTasks() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all project tasks
	True := true
	PageSize := 250
	Page := 1
	tasksOps := &teamwork.GetTasksOps{}
	tasksOps.GetFiles = &True
	tasksOps.IncludeCompletedSubtasks = &True
	tasksOps.IncludeCompletedTasks = &True
	tasksOps.NestSubTasks = "yes"
	tasksOps.PageSize = &PageSize
	tasksOps.Page = &Page
	projectTasks, pages, err := conn.GetProjectTasks("158721", tasksOps)
	if err != nil {
		fmt.Printf("Error getting Project Tasks: %s", err.Error())
	}

	fmt.Println("GetProjectTasks")
	fmt.Println("1. Task List Name:", projectTasks[0].TaskListName)
	fmt.Println("1. Task ID:", projectTasks[0].ID)
	fmt.Println("1. Task Content:", projectTasks[0].Content)
	fmt.Println("1. Task Status:", projectTasks[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetTaskListTasks() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all task list tasks
	True := true
	PageSize := 250
	Page := 1
	tasksOps := &teamwork.GetTasksOps{}
	tasksOps.GetFiles = &True
	tasksOps.IncludeCompletedSubtasks = &True
	tasksOps.IncludeCompletedTasks = &True
	tasksOps.NestSubTasks = "yes"
	tasksOps.PageSize = &PageSize
	tasksOps.Page = &Page
	taskListTasks, pages, err := conn.GetTaskListTasks("704748", tasksOps)
	if err != nil {
		fmt.Printf("Error getting Task List Tasks: %s", err.Error())
	}

	fmt.Println("GetTaskListTasks")
	fmt.Println("1. Task List Name:", taskListTasks[0].TaskListName)
	fmt.Println("1. Task ID:", taskListTasks[0].ID)
	fmt.Println("1. Task Content:", taskListTasks[0].Content)
	fmt.Println("1. Task Status:", taskListTasks[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectTaskLists() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}

	// get all task lists for a project
	projectTaskListsOps := &teamwork.GetProjectTaskListsOps{
		GetOverdueCount:   "yes",
		GetCompletedCount: "yes",
		ShowMilestones:    "1",
	}
	projectTaskLists, pages, err := conn.GetProjectTaskLists("158721", projectTaskListsOps)
	if err != nil {
		fmt.Printf("Error getting Project Task Lists: %s", err.Error())
	}

	fmt.Println("GetProjectTaskLists")
	fmt.Println("1. Task Lists Name:", projectTaskLists[0].Name)
	fmt.Println("1. Task Lists ID:", projectTaskLists[0].ID)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}
