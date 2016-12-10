package teamwork

import (
	"encoding/json"
	"fmt"
	"time"
)

// A list of Tasks.
type Tasks []Task

// The Task structure.
type Task struct {
	Attachments []struct {
		CategoryID   int    `json:"categoryId"`
		CategoryName string `json:"categoryName"`
		FileID       int    `json:"fileId"`
		Filename     string `json:"filename"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Size         string `json:"size"`
		Version      string `json:"version"`
	} `json:"attachments,omitempty"`
	AttachmentsCount          int           `json:"attachments-count"`
	CanComplete               bool          `json:"canComplete"`
	CanEdit                   bool          `json:"canEdit"`
	CanLogTime                bool          `json:"canLogTime"`
	CommentFollowerIds        string        `json:"commentFollowerIds,omitempty"`
	CommentFollowerSummary    string        `json:"commentFollowerSummary,omitempty"`
	CommentsCount             int           `json:"comments-count"`
	CompanyID                 int           `json:"company-id"`
	CompanyName               string        `json:"company-name"`
	Completed                 bool          `json:"completed"`
	CompletedOn               time.Time     `json:"completed_on,omitempty"`
	CompleterFirstname        string        `json:"completer_firstname,omitempty"`
	CompleterID               string        `json:"completer_id,omitempty"`
	CompleterLastname         string        `json:"completer_lastname,omitempty"`
	Content                   string        `json:"content"`
	CreatedOn                 time.Time     `json:"created-on"`
	CreatorAvatarURL          string        `json:"creator-avatar-url"`
	CreatorFirstname          string        `json:"creator-firstname"`
	CreatorID                 int           `json:"creator-id"`
	CreatorLastname           string        `json:"creator-lastname"`
	Description               string        `json:"description"`
	DLM                       int           `json:"DLM"`
	DueDate                   string        `json:"due-date"`
	DueDateBase               string        `json:"due-date-base"`
	EstimatedMinutes          int           `json:"estimated-minutes"`
	HarvestEnabled            bool          `json:"harvest-enabled"`
	HasDependencies           int           `json:"has-dependencies"`
	HasPredecessors           int           `json:"has-predecessors"`
	HasReminders              bool          `json:"has-reminders"`
	HasTickets                bool          `json:"hasTickets"`
	HasUnreadComments         bool          `json:"has-unread-comments"`
	ID                        int           `json:"id"`
	LastChangedOn             time.Time     `json:"last-changed-on"`
	LockdownID                string        `json:"lockdownId"`
	Order                     int           `json:"order"`
	ParentTaskID              string        `json:"parentTaskId"`
	Position                  int           `json:"position"`
	Predecessors              []interface{} `json:"predecessors"`
	Priority                  string        `json:"priority"`
	Private                   int           `json:"private"`
	Progress                  int           `json:"progress"`
	ProjectID                 int           `json:"project-id"`
	ProjectName               string        `json:"project-name"`
	ResponsiblePartyFirstname string        `json:"responsible-party-firstname,omitempty"`
	ResponsiblePartyID        string        `json:"responsible-party-id,omitempty"`
	ResponsiblePartyIds       string        `json:"responsible-party-ids,omitempty"`
	ResponsiblePartyLastname  string        `json:"responsible-party-lastname,omitempty"`
	ResponsiblePartyNames     string        `json:"responsible-party-names,omitempty"`
	ResponsiblePartySummary   string        `json:"responsible-party-summary,omitempty"`
	ResponsiblePartyType      string        `json:"responsible-party-type,omitempty"`
	StartDate                 string        `json:"start-date"`
	Status                    string        `json:"status"`
	SubTasks                  []Task        `json:"subTasks,omitempty"`
	TasklistIsTemplate        bool          `json:"tasklist-isTemplate"`
	TasklistLockdownID        string        `json:"tasklist-lockdownId"`
	TasklistPrivate           bool          `json:"tasklist-private"`
	Tags                      []struct {
		Color string `json:"color"`
		ID    int    `json:"id"`
		Name  string `json:"name"`
	} `json:"tags,omitempty"`
	TimeIsLogged          string `json:"timeIsLogged"`
	TodoListID            int    `json:"todo-list-id"`
	TodoListName          string `json:"todo-list-name"`
	UserFollowingChanges  bool   `json:"userFollowingChanges"`
	UserFollowingComments bool   `json:"userFollowingComments"`
	ViewEstimatedTime     bool   `json:"viewEstimatedTime"`
}

// GetTasksOps is used to generate the query params for the
// GetTasks API call.
type GetTasksOps struct {
	// Query tasks based on these values.
	//
	// Will only return tasks that have been completed after a specified date.
	// Format: "YYYYMMDDHHMMSS"
	CompletedAfterDate string `param:"completedAfterDate"`
	// Will only return tasks that have been completed before a specified date.
	// Format: "YYYYMMDDHHMMSS"
	CompletedBeforeDate string `param:"completedBeforeDate"`
	// For requesting tasks made by a specific person or people.
	// Format is a comma separated list of Person IDs.
	CreatorIDs string `param:"creator-ids"`
	// Must be used in conjunction with StartDate.
	// Format: "YYYYMMDD"
	EndDate string `param:"enddate"`
	// Tasks can be filtered by due dates using the following.
	// Valid Input: "all", "anytime", "overdue", "today", "tomorrow", "thisweek",
	//   "within7", "within14", "within30", "within365", "nodate", "nostartdate", "completed"
	// Default: "anytime"
	Filter string `param:"filter"`
	// Files attached to tasks can be returned within the task object by setting this parameter to true.
	// Valid Input: true, false
	// Default: false
	GetFiles *bool `param:"getFiles"`
	// Subtasks can be excluded from the results by adding this parameter with no as the value.
	// Valid Input: "no", "yes"
	// Default: "yes"
	GetSubTasks string `param:"getSubTasks"`
	// When using the filter option, you can choose to include start dates matching the filtering critera
	// by passing this parameter as true. By default, only due dates are checked against the filter.
	// Valid Input: true, false
	// Default: false
	IgnoreStartDates *bool `param:"ignore-start-dates"`
	// Extra tasks that can be included with the filter option.
	// Valid Input: "nodate", "nostartdate", "noduedate", "overdue"
	Include string `param:"include"`
	// Sub-tasks that have been marked as completed can be shown by setting this parameter to true
	// if you have requested to include sub-tasks
	// Valid Input: true, false
	// Default: false
	IncludeCompletedSubtasks *bool `param:"includeCompletedSubtasks"`
	// Tasks that have been marked as completed can be shown by setting this parameter to true.
	// Valid Input: true, false
	// Default: false
	IncludeCompletedTasks *bool `param:"includeCompletedTasks"`
	// When using the filter option with any of the following options; within7,within14,within30,within365.
	// You can choose to exclude deadlines for today by passing this parameter as false.
	// Valid Input: true, false
	// Default: true
	IncludeToday *bool `param:"includeToday"`
	// Subtasks can be nested within the parent task object by adding this paramter with yes as the value.
	// Default: "no"
	NestSubTasks string `param:"nestSubTasks"`
	// Optionally, you can set the page from which to start retrieving results.
	// This is usually used in conjunction with the parameter PageSize.
	// Default: 1
	Page *int `param:"page"`
	// The amount of tasks returned can be limited using this parameter.
	// Normally used in conjunction with the Page parameter.
	// Default: 250
	PageSize *int `param:"pageSize"`
	// Tasks can be filtered by the person/people a task is assigned to.
	// Details:
	// "-1" would return all tasks with an assigned person.
	// "0" would return all tasks with no assignment.
	// "32" would return tasks assigned to user 32.
	// "32,55" would return tasks assigned to users 32 and/or 55 etc.
	ResponsiblePartyIDs string `param:"responsible-party-ids"`
	// Tasks that have been deleted can be shown by setting this parameter to "yes".
	// Valid Input: "no", "yes"
	// Default: "no"
	ShowDeleted string `param:"showDeleted"`
	// Tasks can be sorted by the following options.
	// Valid Input: "duedate", "startdate", "dateadded", "priority", "project", "company"
	Sort string `param:"sort"`
	// Tasks within a range of dates can be returned by setting a StartDate and EndDate.
	// Format: "YYYYMMDD".
	StartDate string `param:"startdate"`
	// A comma separated list of Tag IDs to filter tasks on.
	TagIDs string `param:"tag-ids"`
	// Will only return tasks that have been updated after a specified date.
	// Format: "YYYYMMDDHHMMSS"
	UpdatedAfterDate string `param:"updatedAfterDate"`
}

// GetTasks gets all the time entries available according to the specified
// GetTasksOps which are passed in.
//
// ref: http://developer.teamwork.com/timetracking#retrieve_all_time
func (conn *Connection) GetTasks(ops *GetTasksOps) (Tasks, Pages, error) {
	tasks := make(Tasks, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%stasks.json%s", conn.Account.Url, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return tasks, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*Tasks `json:"todo-items"`
	}{&tasks})
	if err != nil {
		return tasks, *pages, err
	}

	return tasks, *pages, nil
}
