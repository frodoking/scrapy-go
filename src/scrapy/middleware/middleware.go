package middleware

import "scrapy/http/request"
import "scrapy/spiders"

type MiddlewareManager struct {
	middlewares []MiddlewareManager
	methods map[string]string
}


func (m *MiddlewareManager) download(downloadFunc interface{}, request *request.Request, spider *spiders.Spider) {

}