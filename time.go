package teamwork

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// TimeEntries is a list of TimeEntry
type TimeEntries []TimeEntry

// TimeEntry is a description of a time entry
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
	UserID int `param:"userId"`
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
	// A page contains 100 entries, but you can increase the page size to a
	// maximum of 500 entries by using 'pageSize' parameter.
	PageSize string `param:"pageSize"`
}

// CreateTimeEntryOps is used to generate the query params for the
// CreateTimeEntry* API call.
type CreateTimeEntryOps struct {
	// Description of time entry
	Description string `json:"description"`
	// ID of the user for the entry
	PersonID string `json:"person-id"`
	// Start Date of the time entry in YYYYMMDD format
	Date string `json:"date"`
	// Start Time of the time entry in HH:MM:SS format
	Time string `json:"time"`
	// Hours logged for the time entry
	Hours string `json:"hours"`
	// Minutes logged for the time entry
	Minutes string `json:"minutes"`
	// billable flag
	// Valid Input: true, false
	IsBillable string `json:"isbillable,omitempty"`
	// task associated with this entry
	TaskID string `json:"task-id,omitempty"`
}

// CreateTimeEntryResponse captures the response returned from a create time entry action
type CreateTimeEntryResponse struct {
	ID     int    `json:"timeLogId"`
	Status string `json:"STATUS"`
}

// DeleteTimeEntryResponse captures the response returned from an update time entry action
type DeleteTimeEntryResponse struct {
	Status string `json:"STATUS"`
}

// GetTimeEntries gets all the time entries available according to the specified
// GetTimeEntriesOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#retrieve_all_time
func (conn *Connection) GetTimeEntries(ops *GetTimeEntriesOps) (TimeEntries, Pages, error) {
	timeEntries := make(TimeEntries, 0)
	pages := &Pages{}
	params := buildParams(ops)
	method := "GET"
	url := fmt.Sprintf("%stime_entries.json%s", conn.Account.Url, params)
	reader, headers, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return timeEntries, *pages, err
	}
	// data, _ := ioutil.ReadAll(reader)
	// fmt.Printf(string(data))
	getHeaders(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TimeEntries `json:"time-entries"`
	}{&timeEntries})
	if err != nil {
		return timeEntries, *pages, err
	}

	return timeEntries, *pages, nil
}

// GetProjectTimeEntries gets all the time entries available for a specific project
// according to the specified GetTimeEntriesOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#retrieve_all_time
func (conn *Connection) GetProjectTimeEntries(id string, ops *GetTimeEntriesOps) (TimeEntries, Pages, error) {
	timeEntries := make(TimeEntries, 0)
	pages := &Pages{}
	params := buildParams(ops)
	method := "GET"
	url := fmt.Sprintf("%sprojects/%s/time_entries.json%s", conn.Account.Url, id, params)
	reader, headers, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return timeEntries, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	getHeaders(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TimeEntries `json:"time-entries"`
	}{&timeEntries})
	if err != nil {
		return timeEntries, *pages, err
	}

	return timeEntries, *pages, nil
}

// CreateTimeEntryForProject creates a time entry for a project
// according to the specified CreateTimeEntryOps which are passed in
//
// ref: https://developer.teamwork.com/projects/api-v1/ref/time-tracking/post-projects-id-time-entries-json
func (conn *Connection) CreateTimeEntryForProject(projectID string, ops *CreateTimeEntryOps) (*CreateTimeEntryResponse, error) {
	jsonBody, err := json.Marshal(struct {
		TimeEntry *CreateTimeEntryOps `json:"time-entry"`
	}{TimeEntry: ops})
	// fmt.Println(string(jsonBody))
	createResponse := &CreateTimeEntryResponse{}
	method := "POST"
	url := fmt.Sprintf("%sprojects/%s/time_entries.json", conn.Account.Url, projectID)
	reader, _, err := request(conn.ApiToken, method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	// data, _ := ioutil.ReadAll(reader)
	// fmt.Printf(string(data))
	// getHeaders(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*CreateTimeEntryResponse `json:"time-entry"`
	}{createResponse})
	if err != nil {
		return nil, err
	}

	return createResponse, nil
}

// CreateTimeEntryForTask creates a time entry for a task
// according to the specified CreateTimeEntryOps which are passed in
//
// ref: https://developer.teamwork.com/projects/api-v1/ref/time-tracking/post-projects-id-time-entries-json
func (conn *Connection) CreateTimeEntryForTask(taskID string, ops *CreateTimeEntryOps) (*CreateTimeEntryResponse, error) {
	jsonBody, err := json.Marshal(struct {
		TimeEntry *CreateTimeEntryOps `json:"time-entry"`
	}{TimeEntry: ops})
	// fmt.Println(string(jsonBody))
	createResponse := &CreateTimeEntryResponse{}
	method := "POST"
	url := fmt.Sprintf("%stasks/%s/time_entries.json", conn.Account.Url, taskID)
	reader, _, err := request(conn.ApiToken, method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	// data, _ := ioutil.ReadAll(reader)
	// fmt.Printf(string(data))
	// getHeaders(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&createResponse)
	if err != nil {
		return nil, err
	}

	return createResponse, nil
}

// DeleteTimeEntry deletes a specific time entry
//
// ref: https://developer.teamwork.com/projects/api-v1/ref/time-tracking/delete-time-entries-id-json
func (conn *Connection) DeleteTimeEntry(id string) (*DeleteTimeEntryResponse, error) {
	method := "DELETE"
	url := fmt.Sprintf("%stime_entries/%s.json", conn.Account.Url, id)
	reader, _, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return nil, err
	}
	// data, _ := ioutil.ReadAll(reader)
	// fmt.Printf(string(data))
	// getHeaders(headers, pages)
	defer reader.Close()

	deleteResponse := &DeleteTimeEntryResponse{}
	err = json.NewDecoder(reader).Decode(deleteResponse)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v", deleteResponse)

	return deleteResponse, nil
}

// GetTaskTimeEntries gets all the time entries available for a specific task
// according to the specified GetTimeEntriesOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#retrieve_all_time
func (conn *Connection) GetTaskTimeEntries(id string, ops *GetTimeEntriesOps) (TimeEntries, Pages, error) {
	timeEntries := make(TimeEntries, 0)
	pages := &Pages{}
	params := buildParams(ops)
	method := "GET"
	url := fmt.Sprintf("%stasks/%s/time_entries.json%s", conn.Account.Url, id, params)
	reader, headers, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return timeEntries, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	getHeaders(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TimeEntries `json:"time-entries"`
	}{&timeEntries})
	if err != nil {
		return timeEntries, *pages, err
	}

	return timeEntries, *pages, nil
}

// TotalTime over whole account.
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

// ProjectTotalTimes is a list of ProjectTotalTime
type ProjectTotalTimes []ProjectTotalTime

// ProjectTotalTime describes the time spent on a project
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

// ProjectTaskListTotalTimes is a list of ProjectTaskListTotalTime
type ProjectTaskListTotalTimes []ProjectTaskListTotalTime

// ProjectTaskListTotalTime is a description of the TaskList total time
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

// ProjectTaskTotalTimes is a list of ProjectTaskTotalTime
type ProjectTaskTotalTimes []ProjectTaskTotalTime

// ProjectTaskTotalTime is a description of the total time for a Task
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
	totalTime := &TotalTime{}
	params := buildParams(ops)
	method := "GET"
	url := fmt.Sprintf("%stime/total.json%s", conn.Account.Url, params)
	reader, _, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return *totalTime, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*TotalTime `json:"time-totals"`
	}{totalTime})
	if err != nil {
		return *totalTime, err
	}

	return *totalTime, nil
}

// GetProjectTotalTime gets total time for a specific project according to the specified
// GetTotalTimeOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#time_totals
func (conn *Connection) GetProjectTotalTime(id string, ops *GetTotalTimeOps) (ProjectTotalTimes, error) {
	projectTotalTime := make(ProjectTotalTimes, 0)
	params := buildParams(ops)
	method := "GET"
	url := fmt.Sprintf("%sprojects/%s/time/total.json%s", conn.Account.Url, id, params)
	reader, _, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return projectTotalTime, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*ProjectTotalTimes `json:"projects"`
	}{&projectTotalTime})
	if err != nil {
		return projectTotalTime, err
	}

	return projectTotalTime, nil
}

// GetTaskListTotalTime gets total time for a specific task list according to the specified
// GetTotalTimeOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#time_totals
func (conn *Connection) GetTaskListTotalTime(id string, ops *GetTotalTimeOps) (ProjectTaskListTotalTimes, error) {
	taskListTotalTime := make(ProjectTaskListTotalTimes, 0)
	params := buildParams(ops)
	method := "GET"
	url := fmt.Sprintf("%stasklists/%s/time/total.json%s", conn.Account.Url, id, params)
	reader, _, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return taskListTotalTime, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*ProjectTaskListTotalTimes `json:"projects"`
	}{&taskListTotalTime})
	if err != nil {
		return taskListTotalTime, err
	}

	return taskListTotalTime, nil
}

// GetTaskTotalTime gets total time for a specific task according to the specified
// GetTotalTimeOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#time_totals
func (conn *Connection) GetTaskTotalTime(id string, ops *GetTotalTimeOps) (ProjectTaskTotalTimes, error) {
	taskTotalTime := make(ProjectTaskTotalTimes, 0)
	params := buildParams(ops)
	method := "GET"
	url := fmt.Sprintf("%stasks/%s/time/total.json%s", conn.Account.Url, id, params)
	reader, _, err := request(conn.ApiToken, method, url, nil)
	if err != nil {
		return taskTotalTime, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*ProjectTaskTotalTimes `json:"projects"`
	}{&taskTotalTime})
	if err != nil {
		return taskTotalTime, err
	}

	return taskTotalTime, nil
}
