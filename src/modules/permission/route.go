package permission

import (
	"common/route"
)


func init() {
    subrouter := route.Router.PathPrefix("/permission/").Subrouter()

	subrouter.HandleFunc("/{user:[0-9]+}", index).Methods("GET").Name("permission_list")
}
