package response

import (
	"container/list"
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
	classes   map[string]interface{}
	mimeTypes *list.List
}

func NewResponseTypes() *ResponseTypes {
	classes := make(map[string]interface{})
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
			classes[mimeType] = &Response{}
			break
		}
	}
	return &ResponseTypes{classes, list.New()}
}

func (rt *ResponseTypes) FromMimeType(mimeType string) interface{} {
	if mimeType == "" {
		return &Response{}
	}

	if value, ok := rt.classes[mimeType]; ok {
		return value
	}

	return &Response{}
}

func (rt *ResponseTypes) FromContentType(contentType string, contentEncoding string) interface{} {
	mimeType := ""
	return rt.FromMimeType(mimeType)
}

func (rt *ResponseTypes) FromHeaders(headers map[string]interface{}) interface{} {
	return &Response{}
}

func (rt *ResponseTypes) FromBody(body interface{}) interface{} {
	return &Response{}
}
