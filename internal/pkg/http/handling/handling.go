package handling

type Request interface {
	Clean()
	Validate() error
}

// func ReadRequest(r *http.Request, request Request) error {
// }
