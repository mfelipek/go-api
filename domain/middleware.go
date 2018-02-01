package domain

import (
	"context"
	"net/http"
)

type ContextHandlerFunc func(http.ResponseWriter, *http.Request, context.Context)

func (h ContextHandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	h(rw, r, ctx)
}

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (m MiddlewareFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m(rw, r, next)
}

type ContextMiddlewareFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, ctx context.Context)

func (m ContextMiddlewareFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc, ctx context.Context) {
	m(rw, r, next, ctx)
}

type IMiddleware interface {
	Handler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type IContextMiddleware interface {
	Handler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc, ctx context.Context)
}