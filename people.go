package teamwork

import (
	"encoding/json"
	"fmt"
	"time"
)

// A list of People.
type People []Person

type Person struct {
	Address struct {
		City        string `json:"city"`
		Country     string `json:"country"`
		Countrycode string `json:"countrycode"`
		Line1       string `json:"line1"`
		Line2       string `json:"line2"`
		State       string `json:"state"`
		Zipcode     string `json:"zipcode"`
	} `json:"address"`
	AddressCity    string `json:"address-city"`
	AddressCountry string `json:"address-country"`
	AddressLine1   string `json:"address-line-1"`
	AddressLine2   string `json:"address-line-2"`
	AddressState   string `json:"address-state"`
	AddressZip     string `json:"address-zip"`
	Administrator  bool   `json:"administrator"`
	AvatarURL      string `json:"avatar-url"`
	//CompanyID              string    `json:"company-id"`
	CompanyID               string    `json:"companyId"`
	CompanyName             string    `json:"company-name"`
	CreatedAt               time.Time `json:"created-at"`
	Deleted                 bool      `json:"deleted"`
	DocumentEditorInstalled bool      `json:"documentEditorInstalled"` // Person only
	EmailAddress            string    `json:"email-address"`
	EmailAlt1               string    `json:"email-alt-1"`
	EmailAlt2               string    `json:"email-alt-2"`
	EmailAlt3               string    `json:"email-alt-3"`
	FirstName               string    `json:"first-name"`
	HasAccessToNewProjects  bool      `json:"has-access-to-new-projects"`
	HasDeskAccount          bool      `json:"has-desk-account"` // Person only
	ID                      string    `json:"id"`
	ImHandle                string    `json:"im-handle"`
	ImService               string    `json:"im-service"`
	// bool in People and string in Person :(
	//InOwnerCompany          bool      `json:"in-owner-company"`
	IsClockedIn   bool      `json:"isClockedIn"` // Person only
	LastActive    string    `json:"last-active"`
	LastChangedOn time.Time `json:"last-changed-on"`
	//LastLogin              time.Time `json:"last-login"` // fails if ""
	LastName     string `json:"last-name"`
	Localization struct {
		DateFormat            string `json:"dateFormat"`
		DateFormatID          string `json:"dateFormatId"`
		Language              string `json:"language"`
		LanguageCode          string `json:"languageCode"`
		StartOnSunday         bool   `json:"start-on-sunday"`
		TimeFormat            string `json:"timeFormat"`
		TimeFormatID          string `json:"timeFormatId"`
		Timezone              string `json:"timezone"`
		TimezoneID            string `json:"timezoneId"`
		TimezoneJavaRefCode   string `json:"timezoneJavaRefCode"`
		TimezoneUTCOffsetMins string `json:"timezoneUTCOffsetMins"`
	} `json:"localization"`
	LoginCount  string `json:"login-count"`
	Notes       string `json:"notes"`
	OpenID      string `json:"openId"`
	Permissions struct {
		AddFiles                          string `json:"add-files"`
		AddLinks                          string `json:"add-links"`
		AddMessages                       string `json:"add-messages"`
		AddMilestones                     string `json:"add-milestones"`
		AddNotebooks                      string `json:"add-notebooks"`
		AddPeopleToProject                string `json:"add-people-to-project"`
		AddTaskLists                      string `json:"add-taskLists"`
		AddTasks                          string `json:"add-tasks"`
		AddTime                           string `json:"add-time"`
		CanAddProjects                    bool   `json:"can-add-projects"`
		CanBeAssignedToTasksAndMilestones string `json:"can-be-assigned-to-tasks-and-milestones"`
		CanManagePeople                   bool   `json:"can-manage-people"`
		CanReceiveEmail                   string `json:"can-receive-email"`
		EditAllTasks                      string `json:"edit-all-tasks"`
		IsObserving                       string `json:"is-observing"`
		ProjectAdministrator              string `json:"project-administrator"`
		SetPrivacy                        string `json:"set-privacy"`
		ViewAllTimeLogs                   string `json:"view-all-time-logs"`
		ViewEstimatedTime                 string `json:"view-estimated-time"`
		ViewInvoices                      string `json:"view-invoices"`
		ViewLinks                         string `json:"view-links"`
		ViewMessagesAndFiles              string `json:"view-messages-and-files"`
		ViewNotebooks                     string `json:"view-notebooks"`
		ViewRiskRegister                  string `json:"view-risk-register"`
		ViewTasksAndMilestones            string `json:"view-tasks-and-milestones"`
		ViewTime                          string `json:"view-time"`
	} `json:"permissions"`
	PhoneNumberFax         string `json:"phone-number-fax"`
	PhoneNumberHome        string `json:"phone-number-home"`
	PhoneNumberMobile      string `json:"phone-number-mobile"`
	PhoneNumberMobileParts struct {
		CountryCode string `json:"countryCode"`
		Phone       string `json:"phone"`
		Prefix      string `json:"prefix"`
	} `json:"phone-number-mobile-parts"`
	PhoneNumberOffice     string        `json:"phone-number-office"`
	PhoneNumberOfficeExt  string        `json:"phone-number-office-ext"`
	Pid                   string        `json:"pid"`
	PrivateNotes          string        `json:"private-notes"`
	PrivateNotesText      string        `json:"private-notes-text"`
	Profile               string        `json:"profile"`      // Person Only
	ProfileText           string        `json:"profile-text"` // Person Only
	Projects              []string      `json:"projects"`
	SiteOwner             bool          `json:"site-owner"`
	Tags                  []interface{} `json:"tags"`
	TextFormat            string        `json:"textFormat"`
	Title                 string        `json:"title"`
	Twitter               string        `json:"twitter"`
	UserName              string        `json:"user-name"`
	UserType              string        `json:"user-type"`
	UserUUID              string        `json:"userUUID"`
	UseShorthandDurations bool          `json:"useShorthandDurations"`
	UserInvited           string        `json:"user-invited"`
	//UserInvitedDate       time.Time     `json:"user-invited-date"` // fails if ""
	UserInvitedStatus string `json:"user-invited-status"`
}

// GetPeopleOps is used to generate the query params for the
// GetPeople API call.
type GetPeopleOps struct {
	// Query people based on these values.
	//
	// Pass this parameter to check if a user exists by email address.
	EmailAddress string `param:"emailaddress"`
	// Pass this parameter to return private notes for users.
	// Valid Input: true, false
	FullProfile *bool `param:"fullprofile"`
	// Pass this parameter to return the ProjectIds the user is a member of.
	// Valid Input: true, false
	ReturnProjectIds *bool `param:"returnProjectIds"`
}

// GetPeople gets all the people available according to the specified
// GetPeopleOps which are passed in.
//
// ref: http://developer.teamwork.com/people#get_people
func (conn *Connection) GetPeople(ops *GetPeopleOps) (People, Pages, error) {
	people := make(People, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%speople.json%s", conn.Account.Url, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return people, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*People `json:"people"`
	}{&people})
	if err != nil {
		return people, *pages, err
	}

	return people, *pages, nil
}

// GetProjectPeople gets project people available according to the specified
// GetPeopleOps and company id passed in.
//
// ref: http://developer.teamwork.com/people#get_all_people_(w
func (conn *Connection) GetProjectPeople(id string, ops *GetPeopleOps) (People, Pages, error) {
	people := make(People, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%sprojects/%s/people.json%s", conn.Account.Url, id, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return people, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*People `json:"people"`
	}{&people})
	if err != nil {
		return people, *pages, err
	}

	return people, *pages, nil
}

// GetCompanyPeople gets company people available according to the specified
// GetPeopleOps and company id passed in.
//
// ref: http://developer.teamwork.com/people#get_people_(withi
func (conn *Connection) GetCompanyPeople(id string, ops *GetPeopleOps) (People, Pages, error) {
	people := make(People, 0)
	pages := &Pages{}
	params := build_params(ops)
	method := "GET"
	url := fmt.Sprintf("%scompanies/%s/people.json%s", conn.Account.Url, id, params)
	reader, headers, err := request(conn.ApiToken, method, url)
	if err != nil {
		return people, *pages, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	get_headers(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*People `json:"people"`
	}{&people})
	if err != nil {
		return people, *pages, err
	}

	return people, *pages, nil
}

// GetPerson gets a single person based on a person ID.
//
// ref: http://developer.teamwork.com/projectsapi#retrieve_a_single
func (conn *Connection) GetPerson(id string) (Person, error) {
	person := &Person{}
	method := "GET"
	url := fmt.Sprintf("%speople/%s.json", conn.Account.Url, id)
	reader, _, err := request(conn.ApiToken, method, url)
	if err != nil {
		return *person, err
	}
	//data, _ := ioutil.ReadAll(reader)
	//fmt.Printf(string(data))
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*Person `json:"person"`
	}{person})
	if err != nil {
		return *person, err
	}

	return *person, nil
}
