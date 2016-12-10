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
		var param_value string
		param_name := v.Type().Field(i).Tag.Get("param") // get value from struct field tag

		is_pointer := false
		var kind reflect.Kind
		// Handle either strings or pointers
		switch {
		case v.Field(i).Kind() == reflect.Ptr:
			kind = v.Field(i).Elem().Kind()
			is_pointer = true
		case v.Field(i).Kind() == reflect.String:
			param_value = v.Field(i).Interface().(string)
		}

		// handle pointers
		switch {
		case is_pointer && kind == reflect.String:
			if v.Field(i).Interface() != nil {
				param_value = *v.Field(i).Interface().(*string)
			}
		case is_pointer && kind == reflect.Bool:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatBool(*v.Field(i).Interface().(*bool))
			}
		case is_pointer && kind == reflect.Int:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatInt(int64(*v.Field(i).Interface().(*int)), 10)
			}
		case is_pointer && kind == reflect.Int8:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatInt(int64(*v.Field(i).Interface().(*int8)), 10)
			}
		case is_pointer && kind == reflect.Int16:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatInt(int64(*v.Field(i).Interface().(*int16)), 10)
			}
		case is_pointer && kind == reflect.Int32:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatInt(int64(*v.Field(i).Interface().(*int32)), 10)
			}
		case is_pointer && kind == reflect.Int64:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatInt(*v.Field(i).Interface().(*int64), 10)
			}
		case is_pointer && kind == reflect.Float32:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatFloat(float64(*v.Field(i).Interface().(*float32)), 'f', -1, 64)
			}
		case is_pointer && kind == reflect.Float64:
			if v.Field(i).Interface() != nil {
				param_value = strconv.FormatFloat(*v.Field(i).Interface().(*float64), 'f', -1, 64)
			}
		}
		if param_name != "" && param_value != "" { // make sure we have what we need to set a param
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
// ref: https://gist.github.com/drewolson/4771479
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
