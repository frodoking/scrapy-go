package middleware

import (
	"scrapy/middleware"
)

type CookiesMiddleware struct {
	*middleware.MiddlewareManager
}