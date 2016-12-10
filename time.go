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
	TicketID            string        `json:"ticket-id"`
	ToDoItemID          string        `json:"todo-item-id"`
	ToDoItemName        string        `json:"todo-item-name"`
	ToDoListID          string        `json:"todo-list-id"`
	ToDoListName        string        `json:"todo-list-name"`
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

// GetTaskTimeEntries gets all the time entries available for a specific todo
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
