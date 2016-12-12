package teamwork

import (
	"encoding/json"
	"fmt"
	"time"
)

// A list of Time Entries.
type TimeEntries []TimeEntry

type TimeEntry struct {
	CanEdit             bool          `json:"canEdit"`
	CompanyID           string        `json:"company-id"`
	CompanyName         string        `json:"company-name"`
	CreatedAt           time.Time     `json:"createdAt"`
	Date                time.Time     `json:"date"`
	DateUserPerspective time.Time     `json:"dateUserPerspective"`
	Description         string        `json:"description"`
	HasStartTime        string        `json:"has-start-time"`
	Hours               string        `json:"hours"`
	ID                  string        `json:"id"`
	InvoiceNo           string        `json:"invoiceNo"`
	IsBillable          string        `json:"isbillable"`
	IsBilled            string        `json:"isbilled"`
	Minutes             string        `json:"minutes"`
	ParentTaskID        string        `json:"parentTaskId"`
	ParentTaskName      string        `json:"parentTaskName"`
	PersonFirstName     string        `json:"person-first-name"`
	PersonID            string        `json:"person-id"`
	PersonLastName      string        `json:"person-last-name"`
	ProjectID           string        `json:"project-id"`
	ProjectName         string        `json:"project-name"`
	ProjectStatus       string        `json:"project-status"`
	Tags                []interface{} `json:"tags"`
	TaskEstimatedTime   string        `json:"taskEstimatedTime"`
	TaskIsPrivate       string        `json:"taskIsPrivate"`
	TaskIsSubTask       string        `json:"taskIsSubTask"`
	TaskItemID          string        `json:"todo-item-id"`
	TaskItemName        string        `json:"todo-item-name"`
	TaskListID          string        `json:"todo-list-id"`
	TaskListName        string        `json:"todo-list-name"`
	TicketID            string        `json:"ticket-id"`
	UpdatedDate         time.Time     `json:"updated-date"`
}

// GetTimeEntriesOps is used to generate the query params for the
// GetTimeEntries API call.
type GetTimeEntriesOps struct {
	// Query time entries based on these values.
	//
	// The page to start retrieving entries from (eg: Page = 1 gives records 1 - 100, Page = 2 gives records 101-201 etc)
	Page *int `param:"page"`
	// The start date to retrieve from.
	// Format: "YYYYMMDD"
	FromDate string `param:"fromdate"`
	// The start time only if fromdate is passed.
	// Format: "HH:MM"
	FromTime string `param:"fromtime"`
	// The end date to retrieve to.
	// Format: "YYYYMMDD"
	ToDate string `param:"todate"`
	// The end time only if todate is passed.
	// Format: "HH:MM"
	ToTime string `param:"totime"`
	// Valid Input: "date", "user", "task", "tasklist", "project", "company", "dateupdated"
	// Default: "date"
	SortBy string `param:"sortby"`
	// The order to sort the returned data.
	// Valid Input: "ASC", "DESC"
	SortOrder string `param:"sortorder"`
	// Return time logs for a specific user only
	UserID string `param:"userId"`
	// Filter the Time Entries to those that are Billable or Not Billable.
	// Valid Input: "billable", "nonbillable"
	BillableType string `param:"billableType"`
	// filter the time entries to those that have been Invoiced or not.
	// Valid Input: "invoiced", "noninvoiced"
	InvoicedType string `param:"invoicedType"`
	// Filter the time entries to those in Active projects, Archived projects or All projects.
	// Valid input: "all", "active", "archived"
	ProjectType string `param:"projectType"`
	// Filter time entries to include deleted time sheet entries or not.
	// Valid Input: true, false
	ShowDeleted *bool `param:"showDeleted"`
}

// GetTimeEntries gets all the time entries available according to the specified
// GetTimeEntriesOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#retrieve_all_time
func (conn *Connection) GetTimeEntries(ops *GetTimeEntriesOps) (TimeEntries, Pages, error) {
	time_entries := make(TimeEntries, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%stime_entries.json%s", conn.Account.Url, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return time_entries, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TimeEntries `json:"time-entries"`
	}{&time_entries})
	if err != nil {
		return time_entries, *pages, err
	}

	return time_entries, *pages, nil
}

// GetProjectTimeEntries gets all the time entries available for a specific project
// according to the specified GetTimeEntriesOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#retrieve_all_time
func (conn *Connection) GetProjectTimeEntries(id string, ops *GetTimeEntriesOps) (TimeEntries, Pages, error) {
	time_entries := make(TimeEntries, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%sprojects/%s/time_entries.json%s", conn.Account.Url, id, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return time_entries, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TimeEntries `json:"time-entries"`
	}{&time_entries})
	if err != nil {
		return time_entries, *pages, err
	}

	return time_entries, *pages, nil
}

// GetTaskTimeEntries gets all the time entries available for a specific task
// according to the specified GetTimeEntriesOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#retrieve_all_time
func (conn *Connection) GetTaskTimeEntries(id string, ops *GetTimeEntriesOps) (TimeEntries, Pages, error) {
	time_entries := make(TimeEntries, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%stasks/%s/time_entries.json%s", conn.Account.Url, id, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return time_entries, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TimeEntries `json:"time-entries"`
	}{&time_entries})
	if err != nil {
		return time_entries, *pages, err
	}

	return time_entries, *pages, nil
}

// Total Time over whole account.
type TotalTime struct {
	BillableHoursSum    string `json:"billable-hours-sum"`
	BillableMinsSum     string `json:"billable-mins-sum"`
	BilledHoursSum      string `json:"billed-hours-sum"`
	BilledMinsSum       string `json:"billed-mins-sum"`
	NonBillableHoursSum string `json:"non-billable-hours-sum"`
	NonBillableMinsSum  string `json:"non-billable-mins-sum"`
	NonBilledHoursSum   string `json:"non-billed-hours-sum"`
	NonBilledMinsSum    string `json:"non-billed-mins-sum"`
	TotalHoursSum       string `json:"total-hours-sum"`
	TotalMinsSum        string `json:"total-mins-sum"`
}

type ProjectTotalTimes []ProjectTotalTime

type ProjectTotalTime struct {
	Company struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"company"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	TimeEstimates struct {
		ActiveHoursEstimated                string `json:"active-hours-estimated"`
		ActiveMinsEstimated                 string `json:"active-mins-estimated"`
		CompletedHoursEstimated             string `json:"completed-hours-estimated"`
		CompletedMinsEstimated              string `json:"completed-mins-estimated"`
		FilteredEstimatedHoursSum           string `json:"filtered-estimated-hours-sum"`
		FilteredEstimatedMinsSum            string `json:"filtered-estimated-mins-sum"`
		TotalHoursEstimated                 string `json:"total-hours-estimated"`
		TotalMinsEstimated                  string `json:"total-mins-estimated"`
		TotalWithTimeLoggedEstimatedDecimal string `json:"totalWithTimeLoggedEstimatedDecimal"`
		TotalWithTimeLoggedEstimatedMins    string `json:"totalWithTimeLoggedEstimatedMins"`
	} `json:"time-estimates"`
	TimeTotals struct {
		BillableHoursSum          string `json:"billable-hours-sum"`
		BillableMinsSum           string `json:"billable-mins-sum"`
		BilledHoursSum            string `json:"billed-hours-sum"`
		BilledMinsSum             string `json:"billed-mins-sum"`
		FilteredEstimatedHoursSum string `json:"filtered-estimated-hours-sum"`
		FilteredEstimatedMinsSum  string `json:"filtered-estimated-mins-sum"`
		NonBillableHoursSum       string `json:"non-billable-hours-sum"`
		NonBillableMinsSum        string `json:"non-billable-mins-sum"`
		NonBilledHoursSum         string `json:"non-billed-hours-sum"`
		NonBilledMinsSum          string `json:"non-billed-mins-sum"`
		TotalHoursSum             string `json:"total-hours-sum"`
		TotalMinsSum              string `json:"total-mins-sum"`
	} `json:"time-totals"`
}

type ProjectTaskListTotalTimes []ProjectTaskListTotalTime

type ProjectTaskListTotalTime struct {
	Company struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"company"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	TaskList struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		TimeEstimates struct {
			ActiveHoursEstimated                string `json:"active-hours-estimated"`
			ActiveMinsEstimated                 string `json:"active-mins-estimated"`
			CompletedHoursEstimated             string `json:"completed-hours-estimated"`
			CompletedMinsEstimated              string `json:"completed-mins-estimated"`
			FilteredEstimatedHoursSum           string `json:"filtered-estimated-hours-sum"`
			FilteredEstimatedMinsSum            string `json:"filtered-estimated-mins-sum"`
			TotalHoursEstimated                 string `json:"total-hours-estimated"`
			TotalMinsEstimated                  string `json:"total-mins-estimated"`
			TotalWithTimeLoggedEstimatedDecimal string `json:"totalWithTimeLoggedEstimatedDecimal"`
			TotalWithTimeLoggedEstimatedMins    string `json:"totalWithTimeLoggedEstimatedMins"`
		} `json:"time-estimates"`
		TimeTotals struct {
			BillableHoursSum          string `json:"billable-hours-sum"`
			BillableMinsSum           string `json:"billable-mins-sum"`
			BilledHoursSum            string `json:"billed-hours-sum"`
			BilledMinsSum             string `json:"billed-mins-sum"`
			FilteredEstimatedHoursSum string `json:"filtered-estimated-hours-sum"`
			FilteredEstimatedMinsSum  string `json:"filtered-estimated-mins-sum"`
			NonBillableHoursSum       string `json:"non-billable-hours-sum"`
			NonBillableMinsSum        string `json:"non-billable-mins-sum"`
			NonBilledHoursSum         string `json:"non-billed-hours-sum"`
			NonBilledMinsSum          string `json:"non-billed-mins-sum"`
			TotalHoursSum             string `json:"total-hours-sum"`
			TotalMinsSum              string `json:"total-mins-sum"`
		} `json:"time-totals"`
	} `json:"tasklist"`
}

type ProjectTaskTotalTimes []ProjectTaskTotalTime

type ProjectTaskTotalTime struct {
	Company struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"company"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	TaskList struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Task struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			TimeEstimates struct {
				ActiveHoursEstimated                string `json:"active-hours-estimated"`
				ActiveMinsEstimated                 string `json:"active-mins-estimated"`
				CompletedHoursEstimated             string `json:"completed-hours-estimated"`
				CompletedMinsEstimated              string `json:"completed-mins-estimated"`
				FilteredEstimatedHoursSum           string `json:"filtered-estimated-hours-sum"`
				FilteredEstimatedMinsSum            string `json:"filtered-estimated-mins-sum"`
				TotalHoursEstimated                 string `json:"total-hours-estimated"`
				TotalMinsEstimated                  string `json:"total-mins-estimated"`
				TotalWithTimeLoggedEstimatedDecimal string `json:"totalWithTimeLoggedEstimatedDecimal"`
				TotalWithTimeLoggedEstimatedMins    string `json:"totalWithTimeLoggedEstimatedMins"`
			} `json:"time-estimates"`
			TimeTotals struct {
				BillableHoursSum          string `json:"billable-hours-sum"`
				BillableMinsSum           string `json:"billable-mins-sum"`
				BilledHoursSum            string `json:"billed-hours-sum"`
				BilledMinsSum             string `json:"billed-mins-sum"`
				FilteredEstimatedHoursSum string `json:"filtered-estimated-hours-sum"`
				FilteredEstimatedMinsSum  string `json:"filtered-estimated-mins-sum"`
				NonBillableHoursSum       string `json:"non-billable-hours-sum"`
				NonBillableMinsSum        string `json:"non-billable-mins-sum"`
				NonBilledHoursSum         string `json:"non-billed-hours-sum"`
				NonBilledMinsSum          string `json:"non-billed-mins-sum"`
				TotalHoursSum             string `json:"total-hours-sum"`
				TotalMinsSum              string `json:"total-mins-sum"`
			} `json:"time-totals"`
		} `json:"task"`
	} `json:"tasklist"`
}

// GetTotalTimeOps is used to generate the query params for the
// GetTotalTime API call.
type GetTotalTimeOps struct {
	// Query time entries based on these values.
	//
	// Only show totals for userId passed
	// Default: 0 (All Users)
	UserId *int `param:"userId"`
	// Only show totals from a specific date
	// Format: "YYYYMMDD"
	FromDate string `param:"fromDate"`
	// Only show totals up to a specific date
	// Format: "YYYYMMDD"
	ToDate string `param:"toDate"`
	// Only show totals from a specific time in conjunction with fromDate
	// Format: "HH:MM"
	FromTime string `param:"fromTime"`
	// Only show totals up to a specific time in conjunction with toDate
	// Format: "HH:MM"
	ToTime string `param:"toTime"`
	// Filter the time totals to those in Active projects, Archived projects or All projects.
	// Valid Input: "all", "active", "archived"
	ProjectType string `param:"projectType"`
}

// GetTotalTime gets the total time across all projects according to the specified
// GetTotalTimeOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#time_totals
func (conn *Connection) GetTotalTime(ops *GetTotalTimeOps) (TotalTime, error) {
	total_time := &TotalTime{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%stime/total.json%s", conn.Account.Url, params)
	reader, _, err := request(conn.ApiToken, method, url)
	if err != nil {
		return *total_time, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TotalTime `json:"time-totals"`
	}{total_time})
	if err != nil {
		return *total_time, err
	}

	return *total_time, nil
}

// GetProjectTotalTime gets total time for a specific project according to the specified
// GetTotalTimeOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#time_totals
func (conn *Connection) GetProjectTotalTime(id string, ops *GetTotalTimeOps) (ProjectTotalTimes, error) {
	project_total_time := make(ProjectTotalTimes, 0)
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%sprojects/%s/time/total.json%s", conn.Account.Url, id, params)
	reader, _, err := request(conn.ApiToken, method, url)
	if err != nil {
		return project_total_time, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*ProjectTotalTimes `json:"projects"`
	}{&project_total_time})
	if err != nil {
		return project_total_time, err
	}

	return project_total_time, nil
}

// GetTaskListTotalTime gets total time for a specific task list according to the specified
// GetTotalTimeOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#time_totals
func (conn *Connection) GetTaskListTotalTime(id string, ops *GetTotalTimeOps) (ProjectTaskListTotalTimes, error) {
	task_list_total_time := make(ProjectTaskListTotalTimes, 0)
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%stasklists/%s/time/total.json%s", conn.Account.Url, id, params)
	reader, _, err := request(conn.ApiToken, method, url)
	if err != nil {
		return task_list_total_time, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*ProjectTaskListTotalTimes `json:"projects"`
	}{&task_list_total_time})
	if err != nil {
		return task_list_total_time, err
	}

	return task_list_total_time, nil
}

// GetTaskTotalTime gets total time for a specific task according to the specified
// GetTotalTimeOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#time_totals
func (conn *Connection) GetTaskTotalTime(id string, ops *GetTotalTimeOps) (ProjectTaskTotalTimes, error) {
	task_total_time := make(ProjectTaskTotalTimes, 0)
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%stasks/%s/time/total.json%s", conn.Account.Url, id, params)
	reader, _, err := request(conn.ApiToken, method, url)
	if err != nil {
		return task_total_time, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*ProjectTaskTotalTimes `json:"projects"`
	}{&task_total_time})
	if err != nil {
		return task_total_time, err
	}

	return task_total_time, nil
}
