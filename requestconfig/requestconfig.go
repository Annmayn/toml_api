package requestconfig

//RequestConfig : Global struct to handle all request body
type RequestConfig struct {
	auth         string
	query        string
	result       string
	attachments  []string
	display      interface{}
	data         string
	validation   []string
	schema       string
	query_params []string
}

//Request : request format
type Request struct {
	auth      string
	formatter string
}

//Response : response format
type Response struct {
	formatter string
}

//Backend : backend format
type Backend struct {
	plugin string
}
