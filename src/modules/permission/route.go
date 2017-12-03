package permission

import (
	"common/route"
)


func init() {
    subrouter := route.Router.PathPrefix("/permission/").Subrouter()

	subrouter.HandleFunc("/sync", sync).Methods("GET").Name("permission_sync")
	subrouter.HandleFunc("/list", list).Methods("GET").Name("permission_list")
	subrouter.HandleFunc("/test1", list).Methods("GET").Name("permission_test1")
}
