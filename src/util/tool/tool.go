package tool

import (
	"strings"
	"github.com/gorilla/mux"
	"github.com/Machiel/slugify"
	"common/route"
)


type RawPermission struct {
	UID string "json:uid"
	Module string "json:module"
	Title string "json:title"
	AsciiTitle string "json:title"
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
			rawPermission.AsciiTitle = ToAscii(routeUidArr[1])
			rawPermission.Title = strings.Title(rawPermission.AsciiTitle)

			parsedRouter = append(parsedRouter, rawPermission)
		}
		return nil
	})
	return parsedRouter
}

func ToSlug (input string) string {
	return slugify.Slugify(input)
}

func ToAscii (input string) string {
	result := slugify.Slugify(input)
	return strings.Replace(result, "-", " ", -1)
}