package teamwork

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

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

func Connect(ApiToken string) (*Connection, error) {
	method := "GET"
	url := "http://authenticate.teamworkpm.net/authenticate.json"
	reader, err := Request(ApiToken, method, url)
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

type Projects struct {
	All []Project `json:"projects"`
}

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

func (conn *Connection) GetProjects() (*Projects, error) {
	method := "GET"
	url := fmt.Sprintf("%sprojects.json", conn.Account.Url)
	reader, err := Request(conn.ApiToken, method, url)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	projects := &Projects{}
	if err := json.NewDecoder(reader).Decode(projects); err != nil {
		return nil, err
	}

	//data, _ := ioutil.ReadAll(reader)
	//log.Printf(string(data))

	return projects, nil
}

func Request(token, method, url string) (io.ReadCloser, error) {
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
