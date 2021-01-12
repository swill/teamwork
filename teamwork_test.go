package teamwork

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

func ExampleConnect() {
	// setup the teamwork connection
	apiToken := "a_teamwork_apiToken"
	conn, err := teamwork.Connect(apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}
}
