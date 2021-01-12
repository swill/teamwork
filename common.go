package teamwork

// Pages provides a way to page requests
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
