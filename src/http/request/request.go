package src

type Request struct {
	url string
	encoding string
	method string
    headers *Headers
	body string
	cookies *Cookies
	meta string
	priority int
	dontFilter bool
    flags []string
}

func meta(request *Request) string {
	return request.meta
}
