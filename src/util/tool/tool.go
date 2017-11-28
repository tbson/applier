package tool

import (
	"strings"
	"github.com/gorilla/mux"
	"common/route"
)


type RawPermission struct {
	UID string "json:uid"
	Module string "json:module"
	Title string "json:title"
}

func ParseRouter () []RawPermission {
	var parsedRouter = make([]RawPermission, 0)
	route.Router.Walk(func (route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeUid := route.GetName();
		if routeUid != "" {
			rawPermission := RawPermission{}
			routeUidArr := strings.Split(routeUid, "_");

			rawPermission.UID = routeUid
			rawPermission.Module = routeUidArr[0]
			rawPermission.Title = routeUidArr[1]

			parsedRouter = append(parsedRouter, rawPermission)
		}
		return nil
	})
	return parsedRouter
}