package teamwork

import (
	"fmt"

	"github.com/swill/teamwork"
)

func ExampleConnection_GetTimeEntries() {
	// get all time entries
	Page := 200
	time_entries_ops := &teamwork.GetTimeEntriesOps{
		Page: &Page,
	}
	time_entries, pages, err := conn.GetTimeEntries(time_entries_ops)
	if err != nil {
		fmt.Printf("Error getting Time Entries: %s", err.Error())
	}

	fmt.Println("GetTimeEntries")
	fmt.Println("1. Time for Company Name:", time_entries[0].CompanyName)
	fmt.Println("1. Time for Project Name:", time_entries[0].ProjectName)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetProjectTimeEntries() {
	// get all project time entries
	project_time_entries_ops := &teamwork.GetTimeEntriesOps{}
	project_time_entries, pages, err := conn.GetProjectTimeEntries("158721", project_time_entries_ops)
	if err != nil {
		fmt.Printf("Error getting Project Time Entries: %s", err.Error())
	}

	fmt.Println("GetProjectTimeEntries")
	fmt.Println("1. Time for Project Name:", project_time_entries[0].ProjectName)
	fmt.Println("1. Time for Project Description:", project_time_entries[0].Description)
	fmt.Println("1. Time for Date:", project_time_entries[0].Date)
	fmt.Println("1. Time in Hours:", project_time_entries[0].Hours)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
}

func ExampleConnection_GetTaskTimeEntries() {
	// get all task time entries
	task_time_entries_ops := &teamwork.GetTimeEntriesOps{}
	task_time_entries, pages, err := conn.GetTaskTimeEntries("4754100", task_time_entries_ops)
	if err != nil {
		fmt.Printf("Error getting Task Time Entries: %s", err.Error())
	}

	fmt.Println("GetTaskTimeEntries")
	fmt.Println("1. Time for Task List:", task_time_entries[0].ToDoListName)
	fmt.Println("1. Time for Task Name:", task_time_entries[0].ToDoItemName)
	fmt.Println("1. Time for Date:", task_time_entries[0].Date)
	fmt.Println("1. Time in Hours:", task_time_entries[0].Hours)
	fmt.Println("on page #:", pages.Page)
	fmt.Println("# of pages:", pages.Pages)
	fmt.Println("# of records:", pages.Records)
}

func ExampleConnection_GetTotalTime() {
	// get total time for an account
	total_time_ops := &teamwork.GetTotalTimeOps{}
	total_time, err := conn.GetTotalTime(total_time_ops)
	if err != nil {
		fmt.Printf("Error getting Total Time: %s", err.Error())
	}

	fmt.Println("GetTotalTime")
	fmt.Println("Total Hours:", total_time.TotalHoursSum)
	fmt.Println("Total Hours Billable:", total_time.BillableHoursSum)
}

func ExampleConnection_GetProjectTotalTime() {
	// get total time for a project
	total_time_ops = &teamwork.GetTotalTimeOps{}
	project_total_time, err := conn.GetProjectTotalTime("158721", total_time_ops)
	if err != nil {
		fmt.Printf("Error getting Project Total Time: %s", err.Error())
	}

	fmt.Println("GetProjectTotalTime")
	fmt.Println("Name:", project_total_time[0].Name)
	fmt.Println("Total Hours:", project_total_time[0].TimeTotals.TotalHoursSum)
	fmt.Println("Total Hours Billable:", project_total_time[0].TimeTotals.BillableHoursSum)
}

func ExampleConnection_GetTaskListTotalTime() {
	// get total time for a task list
	total_time_ops = &teamwork.GetTotalTimeOps{}
	task_list_total_time, err := conn.GetTaskListTotalTime("704748", total_time_ops)
	if err != nil {
		fmt.Printf("Error getting Project Total Time: %s", err.Error())
	}

	fmt.Println("GetTaskListTotalTime")
	fmt.Println("Name:", task_list_total_time[0].TaskList.Name)
	fmt.Println("Total Hours:", task_list_total_time[0].TaskList.TimeTotals.TotalHoursSum)
	fmt.Println("Total Hours Billable:", task_list_total_time[0].TaskList.TimeTotals.BillableHoursSum)
}

func ExampleConnection_GetTaskTotalTime() {
	// get total time for a task
	total_time_ops = &teamwork.GetTotalTimeOps{}
	task_total_time, err := conn.GetTaskTotalTime("4486838", total_time_ops)
	if err != nil {
		fmt.Printf("Error getting Project Total Time: %s", err.Error())
	}

	fmt.Println("GetTaskTotalTime")
	fmt.Println("Name:", task_total_time[0].TaskList.Task.Name)
	fmt.Println("Total Hours:", task_total_time[0].TaskList.Task.TimeTotals.TotalHoursSum)
	fmt.Println("Total Hours Billable:", task_total_time[0].TaskList.Task.TimeTotals.BillableHoursSum)
}
