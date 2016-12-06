package teamwork

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Connection is the established connection to TeamWork.
// Once connected, it populates the details for the Account
// which is making the calls to the API.
type Connection struct {
	Account struct {
		AvatarUrl                  string `json:"avatar-url"`
		CanAddProjects             string `json:"canaddprojects"`
		CanManagePeople            string `json:"canManagePeople"`
		ChatEnabled                bool   `json:"chatEnabled"`
		Code                       string `json:"code"`
		CompanyID                  string `json:"companyid"`
		CompanyName                string `json:"companyname"`
		DateFormat                 string `json:"dateFormat"`
		DateSeperator              string `json:"dateSeperator"`
		DeskEnabled                bool   `json:"deskEnabled"`
		DocumentEditorEnabled      bool   `json:"documentEditorEnabled"`
		FirstName                  string `json:"firstname"`
		ID                         string `json:"id"`
		Lang                       string `json:"lang"`
		LastName                   string `json:"lastname"`
		LikesEnabled               bool   `json:"likesEnabled"`
		Logo                       string `json:"logo"`
		Name                       string `json:"name"`
		PlanID                     string `json:"plan-id"`
		ProjectsEnabled            bool   `json:"projectsEnabled"`
		RequireHttps               bool   `json:"requirehttps"`
		SslEnabled                 bool   `json:"ssl-enabled"`
		StartOnSundays             bool   `json:"startonsundays"`
		TagsEnabled                bool   `json:"tagsEnabled"`
		TagsLockedToAdmins         bool   `json:"tagsLockedToAdmins"`
		TimeFormat                 string `json:"timeFormat"`
		TkoEnabled                 bool   `json:"TKOEnabled"`
		Url                        string `json:"URL"`
		UserID                     string `json:"userId"`
		UserIsAdmin                bool   `json:"userIsAdmin"`
		UserIsMemberOfOwnerCompany string `json:"userIsMemberOfOwnerCompany"`
	} `json:"account"`
	ApiToken string
}

// Connect is the starting point to using the TeamWork API.
// This function returns a Connection which is used to query
// TeamWork via other functions.
func Connect(ApiToken string) (*Connection, error) {
	method := "GET"
	url := "http://authenticate.teamworkpm.net/authenticate.json"
	reader, _, err := request(ApiToken, method, url)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	connection := &Connection{
		ApiToken: ApiToken,
	}
	if err := json.NewDecoder(reader).Decode(connection); err != nil {
		return nil, err
	}
	return connection, nil
}

// The X-Page(s) headers that are returned with queries.
// The struct is populated by the headers when returning
// lists of data from TeamWork.  Use thic concept to
// set a struct from the response headers of the API.
// You only have to specify the `header:"Header-Name"`
// and then use `get_headers(headers, &struct)` to
// populate.
// Currently supports: Int and String
type Pages struct {
	Page    int `header:"X-Page"`
	Pages   int `header:"X-Pages"`
	Records int `header:"X-Records"`
}

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

// ProjectsOps is used to generate the query params for the
// GetProjects API call.
type ProjectsOps struct {
	// Query projects based on these values.
	//
	// The category id to filter by.
	CategoryID string `param:"catId"`
	// The project was created after this date.  Eg: "20100603"
	CreatedAfterDate string `param:"createdAfterDate"`
	// The project was created after this time.  Eg: "15:21"
	CreatedAfterTime string `param:"createdAfterTime"`
	// Output the people include in the project.  Eg: "true"
	IncludePeople string `param:"includePeople"`
	// Order the results by this value. Eg: "name", "companyName", etc...
	OrderBy string `param:"orderby"`
	// A page is 500 results.  Access additional pages.  Eg: "2", "5", etc...
	Page string `param:"page"`
	// The status of the project.
	// Valid values are: ALL, ACTIVE, ARCHIVED, CURRENT, LATE, COMPLETED
	Status string `param:"status"`
	// The project was updated after this date.  Eg: "20100603"
	UpdatedAfterDate string `param:"updatedAfterDate"`
	// The project was updated after this time.  Eg: "15:21"
	UpdatedAfterTime string `param:"updatedAfterTime"`
}

// GetProjects gets all the projects available according to the specified
// ProjectsOps which are passed in.
//
// ref: http://developer.teamwork.com/projectsapi#retrieve_all_proj
func (conn *Connection) GetProjects(ops *ProjectsOps) (Projects, Pages, error) {
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

// ProjectOps is used to generate the query params for the
// GetProject API call.
type ProjectOps struct {
	// Query a project based on these values.
	//
	// Output the people include in the project.  Eg: "true"
	IncludePeople string `param:"includePeople"`
}

// GetProject gets a single project based on a project ID.
//
// ref: http://developer.teamwork.com/projectsapi#retrieve_a_single
func (conn *Connection) GetProject(id string, ops *ProjectOps) (Project, error) {
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

	//data, _ := ioutil.ReadAll(reader)
	//log.Printf(string(data))

	return *project, nil
}

// request is the base level function for calling the TeamWork API.
func request(token, method, url string) (io.ReadCloser, http.Header, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil) // TODO: Add payload to support POST
	if err != nil {
		log.Printf("NewRequest: ", err)
		return nil, nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(token, "notused")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do: ", err)
		return nil, nil, err
	}

	return resp.Body, resp.Header, nil
}

// build_params takes a struct and builds query params based
// on the `param:"param_name"` struct field tags.
//
// ref: https://play.golang.org/p/P9zvVJnMhR
// ref: https://gist.github.com/drewolson/4771479
func build_params(ops interface{}) string {
	pairs := make([]string, 0)
	v := reflect.ValueOf(ops).Elem()
	for i := 0; i < v.NumField(); i++ {
		param_name := v.Type().Field(i).Tag.Get("param")    // get value from struct field tag
		param_value := v.Field(i).Interface()               // value of the field
		if param_name != "" && param_value.(string) != "" { // make sure we have what we need to set a param
			pair := fmt.Sprintf("%s=%s", param_name, param_value)
			pairs = append(pairs, pair) // add to the param pairs array
		}
	}
	if len(pairs) > 0 {
		return fmt.Sprintf("?%s", strings.Join(pairs, "&")) // return the params with the leading '?'
	} else {
		return "" // nothing to send back
	}
}

// get_headers takes the response headers and populates
// a struct of data according to the `header:"HeaderName"`.
// Function currently only supports Int and String field types.
//
// ref: https://play.golang.org/p/P9zvVJnMhR
// ref: https://gist.github.com/drewolson/4771479\
// ref: http://stackoverflow.com/a/6396678/977216
func get_headers(headers http.Header, obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	if v.Kind() == reflect.Struct { // make sure we have a struct
		for i := 0; i < v.NumField(); i++ { // for all fields
			field := v.Field(i)                    // value field.
			if field.IsValid() && field.CanSet() { // is exported and addressable
				header_name := v.Type().Field(i).Tag.Get("header") // get value from struct field tag
				if header_name != "" {                             // make sure the header is set
					header_val := headers.Get(header_name)
					if header_val != "" { // make sure we have a value in the header
						switch {
						case field.Kind() == reflect.Int: // Int struct field type
							h_val, err := strconv.ParseInt(header_val, 10, 64)
							if err != nil {
								log.Printf("Failed to convert header '%s' to a 64 bit Int. \n%s", header_name, err.Error())
								continue
							}
							if !field.OverflowInt(h_val) {
								field.SetInt(h_val)
							}
						case field.Kind() == reflect.String: // String struct field type
							field.SetString(header_val)
						}
					}
				}
			}
		}
	}
}
