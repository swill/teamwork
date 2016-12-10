package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetTasks() {
	// get all tasks
	PageSize := 250
	Page = 2
	tasks_ops := &teamwork.GetTasksOps{}
	tasks_ops.GetFiles = &True
	tasks_ops.IncludeCompletedSubtasks = &True
	tasks_ops.IncludeCompletedTasks = &True
	tasks_ops.NestSubTasks = "yes"
	tasks_ops.PageSize = &PageSize
	tasks_ops.Page = &Page
	tasks, pages, err := conn.GetTasks(tasks_ops)
	if err != nil {
		fmt.Printf("Error getting Tasks: %s", err.Error())
	}

	fmt.Println("1. Task List Name:", tasks[0].TodoListName)
	fmt.Println("1. Task ID:", tasks[0].ID)
	fmt.Println("1. Task Content:", tasks[0].Content)
	fmt.Println("1. Task Status:", tasks[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}
