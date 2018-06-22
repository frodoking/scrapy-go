package response

import (
	"container/list"
	"scrapy/http"
)

var Classes = map[string]string{
	"text/html":                           "HtmlResponse",
	"application/atom+xml":                "XmlResponse",
	"application/rdf+xml":                 "XmlResponse",
	"application/rss+xml":                 "XmlResponse",
	"application/xhtml+xml":               "HtmlResponse",
	"application/vnd.wap.xhtml+xml":       "HtmlResponse",
	"application/xml":                     "XmlResponse",
	"application/json":                    "TextResponse",
	"application/x-json":                  "TextResponse",
	"application/json-amazonui-streaming": "TextResponse",
	"application/javascript":              "TextResponse",
	"application/x-javascript":            "TextResponse",
	"text/xml":                            "XmlResponse",
	"text/*":                              "TextResponse",
}

type ResponseTypes struct {
	classes   map[string]Response
	mimeTypes *list.List
}

func NewResponseTypes() *ResponseTypes {
	classes := make(map[string]Response)
	for mimeType, class := range Classes {
		switch class {
		case "TextResponse":
			classes[mimeType] = &TEXTResponse{}
			break
		case "HtmlResponse":
			classes[mimeType] = &HTMLResponse{}
			break
		case "XmlResponse":
			classes[mimeType] = &XMLResponse{}
			break
		default:
			classes[mimeType] = NewResponse("")
			break
		}
	}
	return &ResponseTypes{classes, list.New()}
}

func (rt *ResponseTypes) FromMimeType(mimeType string) Response {
	if mimeType == "" {
		return NewResponse("")
	}

	if value, ok := rt.classes[mimeType]; ok {
		return value
	}

	return NewResponse("")
}

func (rt *ResponseTypes) FromContentType(contentType string, contentEncoding string) Response {
	mimeType := ""
	return rt.FromMimeType(mimeType)
}

func (rt *ResponseTypes) FromHeaders(headers *http.Headers) Response {
	return NewResponse("")
}

func (rt *ResponseTypes) FromBody(body interface{}) Response {
	return NewResponse("")
}

var ScrapyResponseTypes = NewResponseTypes()
