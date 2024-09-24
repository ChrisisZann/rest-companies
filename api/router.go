package main

import (
	"net/http"
	"xm-companies/events"
)

func (api *api) router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.rootHandler)

	mux.HandleFunc("/company", api.company)
	mux.Handle("/auth-company", api.Auth(http.HandlerFunc(api.auth_company)))

	mux.HandleFunc("/user", api.user)
	mux.HandleFunc("/login", api.login)

	mux.Handle("/auth-login", api.Auth(http.HandlerFunc(api.ProtectedEndpoint)))

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		events.ServeWS(api.hub, w, r)
	})

	return mux

}
