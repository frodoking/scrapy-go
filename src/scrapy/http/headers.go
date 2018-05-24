package http

type Headers struct {
	keys     map[string]interface{}
	encoding string
}

func NewHeaders(keys map[string]interface{}) *Headers {
	headers := &Headers{}
	headers.encoding = "utf-8"
	headers.keys = keys
	return headers
}
