package models

import (
	"context"
	"net/http"
	"github.com/google/uuid"
)

func RequestMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd := Request{}
		rd.From = r.Header.Get("Request")
		rd.TraceId = r.Header.Get("Trace")
		if rd.TraceId == "" {
			rd.TraceId = uuid.New().String()
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "Request", rd)))
	})
}

func CacheMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd := Request{}
		rd.From = r.Header.Get("Request")
		rd.TraceId = r.Header.Get("Trace")
		if rd.TraceId == "" {
			rd.TraceId = uuid.New().String()
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "Request", rd)))
	})
}

type Request struct {
	TraceId string
	From string
}
