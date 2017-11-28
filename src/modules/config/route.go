package config

import (
	"common/route"
)


func init() {
    subrouter := route.Router.PathPrefix("/config/").Subrouter()

	subrouter.HandleFunc("/{user:[0-9]+}", index).Methods("GET").Name("config_list")
}
