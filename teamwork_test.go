package teamwork

import (
	"fmt"
	"os"

	"github.com/swill/teamwork"
)

func ExampleConnect() {
	// setup the teamwork connection
	api_token := "the_teamwork_api_token"
	conn, err := teamwork.Connect(api_token)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}
}
