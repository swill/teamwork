package teamwork

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
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
	reader, err := request(ApiToken, method, url)
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

// A list of Projects,
type Projects []Project

// The Project structure.
type Project struct {
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
	Notifyeveryone       bool      `json:"notifyeveryone"`
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
func (conn *Connection) GetProjects(ops *ProjectsOps) (Projects, error) {
	params := build_params(ops)
	fmt.Println("params", params)
	method := "GET"
	url := fmt.Sprintf("%sprojects.json%s", conn.Account.Url, params)
	reader, err := request(conn.ApiToken, method, url)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	projects := make(Projects, 0)
	err = json.NewDecoder(reader).Decode(&struct {
		*Projects `json:"projects"`
	}{&projects})
	if err != nil {
		return nil, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//log.Printf(string(data))

	return projects, nil
}

// request is the base level function for calling the TeamWork API.
func request(token, method, url string) (io.ReadCloser, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("NewRequest: ", err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(token, "notused")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do: ", err)
		return nil, err
	}

	return resp.Body, nil
}

// build_params takes and struct and build query params based
// on the `param:"param_name"` struct field tags.
//
// ref: https://play.golang.org/p/P9zvVJnMhR
// ref: https://gist.github.com/drewolson/4771479
func build_params(ops interface{}) string {
	pairs := make([]string, 0)
	v := reflect.ValueOf(ops).Elem()
	for i := 0; i < v.NumField(); i++ {
		param_name := v.Type().Field(i).Tag.Get("param")
		param_value := v.Field(i).Interface()
		if param_name != "" && param_value.(string) != "" {
			pair := fmt.Sprintf("%s=%s", param_name, param_value)
			pairs = append(pairs, pair)
		}
	}
	if len(pairs) > 0 {
		return fmt.Sprintf("?%s", strings.Join(pairs, "&"))
	} else {
		return ""
	}
}
