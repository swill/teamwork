package teamwork

import (
	"encoding/json"
	"fmt"
	"time"
)

// A list of Projects.
type Projects []Project

// The Project structure.
type Project struct {
	ActivePages struct {
		Billing      string `json:"billing"`
		Files        string `json:"files"`
		Links        string `json:"links"`
		Messages     string `json:"messages"`
		Milestones   string `json:"milestones"`
		Notebooks    string `json:"notebooks"`
		RiskRegister string `json:"riskRegister"`
		Tasks        string `json:"tasks"`
		Time         string `json:"time"`
	} `json:"active-pages"`
	Announcement     string `json:"announcement"`
	AnnouncementHTML string `json:"announcementHTML"`
	Category         struct {
		Color string `json:"color"`
		ID    string `json:"id"`
		Name  string `json:"name"`
	} `json:"category"`
	Company struct {
		ID      string `json:"id"`
		IsOwner string `json:"is-owner"`
		Name    string `json:"name"`
	} `json:"company"`
	CreatedOn      time.Time `json:"created-on"`
	DefaultPrivacy string    `json:"defaultPrivacy"`
	Defaults       struct {
		Privacy string `json:"privacy"`
	} `json:"defaults"`
	Description          string    `json:"description"`
	EndDate              string    `json:"endDate"`
	FilesAutoNewVersion  bool      `json:"filesAutoNewVersion"`
	HarvestTimersEnabled bool      `json:"harvest-timers-enabled"`
	ID                   string    `json:"id"`
	IsProjectAdmin       bool      `json:"isProjectAdmin"`
	LastChangedOn        time.Time `json:"last-changed-on"`
	Logo                 string    `json:"logo"`
	LogoFromCompany      bool      `json:"logoFromCompany,omitempty"`
	Name                 string    `json:"name"`
	NotifyEveryone       bool      `json:"notifyeveryone"`
	People               []string  `json:"people"`
	PrivacyEnabled       bool      `json:"privacyEnabled"`
	ReplyByEmailEnabled  bool      `json:"replyByEmailEnabled"`
	ShowAnnouncement     bool      `json:"show-announcement"`
	Starred              bool      `json:"starred"`
	StartDate            string    `json:"startDate"`
	StartPage            string    `json:"start-page"`
	Status               string    `json:"status"`
	SubStatus            string    `json:"subStatus"`
	Tags                 []struct {
		Color string `json:"color"`
		ID    string `json:"id"`
		Name  string `json:"name"`
	} `json:"tags"`
}

// GetProjectsOps is used to generate the query params for the
// GetProjects API call.
type GetProjectsOps struct {
	// Query projects based on these values.
	//
	// The category id to filter by.
	CategoryID string `param:"catId"`
	// The project was created after this date.  Eg: "20100603"
	CreatedAfterDate string `param:"createdAfterDate"`
	// The project was created after this time.  Eg: "15:21"
	CreatedAfterTime string `param:"createdAfterTime"`
	// Output the people include in the project.  Eg: true
	IncludePeople bool `param:"includePeople"`
	// Order the results by this value. Eg: "name", "companyName", etc...
	OrderBy string `param:"orderby"`
	// A page is 500 results.  Access additional pages.  Eg: 2, etc...
	Page int `param:"page"`
	// The status of the project.
	// Valid values are: ALL, ACTIVE, ARCHIVED, CURRENT, LATE, COMPLETED
	Status string `param:"status"`
	// The project was updated after this date.  Eg: "20100603"
	UpdatedAfterDate string `param:"updatedAfterDate"`
	// The project was updated after this time.  Eg: "15:21"
	UpdatedAfterTime string `param:"updatedAfterTime"`
}

// GetProjects gets all the projects available according to the specified
// GetProjectsOps which are passed in.
//
// ref: http://developer.teamwork.com/projectsapi#retrieve_all_proj
func (conn *Connection) GetProjects(ops *GetProjectsOps) (Projects, Pages, error) {
	projects := make(Projects, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%sprojects.json%s", conn.Account.Url, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return projects, *pages, err
	}
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*Projects `json:"projects"`
	}{&projects})
	if err != nil {
		return projects, *pages, err
	}

	return projects, *pages, nil
}

// GetProjectOps is used to generate the query params for the
// GetProject API call.
type GetProjectOps struct {
	// Query a project based on these values.
	//
	// Output the people include in the project.  Eg: "true"
	IncludePeople string `param:"includePeople"`
}

// GetProject gets a single project based on a project ID.
//
// ref: http://developer.teamwork.com/projectsapi#retrieve_a_single
func (conn *Connection) GetProject(id string, ops *GetProjectOps) (Project, error) {
	project := &Project{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%sprojects/%s.json%s", conn.Account.Url, id, params)
	reader, _, err := request(conn.ApiToken, method, url)
	if err != nil {
		return *project, err
	}
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*Project `json:"project"`
	}{project})
	if err != nil {
		return *project, err
	}

	return *project, nil
}
