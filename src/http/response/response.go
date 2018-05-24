package src

type Response struct {
	url string
    headers *Headers
	status int
	body []byte
	request *Request
    flags []string
}

func Copy(response *Response)  {
	
}

func Replace(response *Response)  {
	
}

func Xpath(response *Response) interface{} {
	return nil
}

func Follow(response *Response) *Request {
	return nil
}
