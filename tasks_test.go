package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetTasks() {
	// get all tasks
	PageSize := 250
	Page := 2
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

	fmt.Println("GetTasks")
	fmt.Println("1. Task List Name:", tasks[0].TodoListName)
	fmt.Println("1. Task ID:", tasks[0].ID)
	fmt.Println("1. Task Content:", tasks[0].Content)
	fmt.Println("1. Task Status:", tasks[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectTasks() {
	// get all project tasks
	PageSize := 250
	Page := 1
	tasks_ops = &teamwork.GetTasksOps{}
	tasks_ops.GetFiles = &True
	tasks_ops.IncludeCompletedSubtasks = &True
	tasks_ops.IncludeCompletedTasks = &True
	tasks_ops.NestSubTasks = "yes"
	tasks_ops.PageSize = &PageSize
	tasks_ops.Page = &Page
	project_tasks, pages, err := conn.GetProjectTasks("158721", tasks_ops)
	if err != nil {
		fmt.Printf("Error getting Project Tasks: %s", err.Error())
	}

	fmt.Println("GetProjectTasks")
	fmt.Println("1. Task List Name:", project_tasks[0].TaskListName)
	fmt.Println("1. Task ID:", project_tasks[0].ID)
	fmt.Println("1. Task Content:", project_tasks[0].Content)
	fmt.Println("1. Task Status:", project_tasks[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetTaskListTasks() {
	// get all task list tasks
	PageSize := 250
	Page := 1
	tasks_ops = &teamwork.GetTasksOps{}
	tasks_ops.GetFiles = &True
	tasks_ops.IncludeCompletedSubtasks = &True
	tasks_ops.IncludeCompletedTasks = &True
	tasks_ops.NestSubTasks = "yes"
	tasks_ops.PageSize = &PageSize
	tasks_ops.Page = &Page
	task_list_tasks, pages, err := conn.GetTaskListTasks("704748", tasks_ops)
	if err != nil {
		fmt.Printf("Error getting Task List Tasks: %s", err.Error())
	}

	fmt.Println("GetTaskListTasks")
	fmt.Println("1. Task List Name:", task_list_tasks[0].TaskListName)
	fmt.Println("1. Task ID:", task_list_tasks[0].ID)
	fmt.Println("1. Task Content:", task_list_tasks[0].Content)
	fmt.Println("1. Task Status:", task_list_tasks[0].Status)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectTaskLists() {
	// get all task lists for a project
	project_task_lists_ops := &teamwork.GetProjectTaskListsOps{
		GetOverdueCount:   "yes",
		GetCompletedCount: "yes",
		ShowMilestones:    "1",
	}
	project_task_lists, pages, err := conn.GetProjectTaskLists("158721", project_task_lists_ops)
	if err != nil {
		fmt.Printf("Error getting Project Task Lists: %s", err.Error())
	}

	fmt.Println("GetProjectTaskLists")
	fmt.Println("1. Task Lists Name:", project_task_lists[0].Name)
	fmt.Println("1. Task Lists ID:", project_task_lists[0].ID)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}
