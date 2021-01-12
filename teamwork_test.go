package teamwork_test

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

func ExampleConnect() {
	// setup the teamwork connection
	baseURL := "a teamwork baseURL"
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v", conn) // this is just so the go linter doesn't complain about "conn defined and not used"
}
