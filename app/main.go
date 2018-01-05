package main

import (
	"app/user"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

func main() {
	var svc user.UserService
	svc = user.NewService()
	svc = user.NewLoggingMiddleware(svc)
	svc = user.NewInstrumentingMiddleware(svc)
	http.Handle("/get-user", user.MakeGetUserHttpHandler(svc))
	http.Handle("/create-user", user.MakeCreateUserHttpHandler(svc))
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":5000", nil)
}
