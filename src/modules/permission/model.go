package permission

import (
	// "fmt"
	// "encoding/json"
	"log"
	"strings"
	"github.com/gorilla/mux"
	"common/route"
	"common/db"
	"util/tool"
)

type RawPermission struct {
	UID string "json:uid"
	Module string "json:module"
	Title string "json:title"
	AsciiTitle string "json:title"
}

type Permission struct {
	ID int "json:id"
	UID string "json:uid"
	Module string "json:module"
	Title string "json:title"
	AsciiTitle string "json:title"
	CreatedAt string "json:created_at"
	UpdatedAt string "json:updated_at"
}

func ParseRouter () []*RawPermission {
	var parsedRouter = make([]*RawPermission, 0)
	route.Router.Walk(func (route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeUid := route.GetName();
		if routeUid != "" {
			rawPermission := &RawPermission{}
			routeUidArr := strings.Split(routeUid, "_");

			rawPermission.UID = routeUid
			rawPermission.Module = routeUidArr[0]
			rawPermission.AsciiTitle = tool.ToAscii(routeUidArr[1])
			rawPermission.Title = strings.Title(rawPermission.AsciiTitle)

			parsedRouter = append(parsedRouter, rawPermission)
		}
		return nil
	})
	return parsedRouter
}

func List() ([]*Permission, error) {
	var result = make([]*Permission, 0)

	rows, err := db.Db.Query("SELECT id, uid, module, title, ascii_title, created_at, updated_at FROM permission")
    if err != nil {
        return result, err
    }
    defer rows.Close()

	for rows.Next() {
		item := Permission{}
		err := rows.Scan(&item.ID, &item.UID, &item.Module, &item.Title, &item.AsciiTitle, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func Sync() ([]*Permission, error) {
	listRoute := ParseRouter();
	for _, item := range listRoute {
		insertStatement := `
			INSERT INTO permission (uid, module, title, ascii_title)
			SELECT CAST($1 AS VARCHAR), $2, $3, $4
			WHERE NOT EXISTS (
				SELECT 1 FROM permission WHERE uid = $1
			)
		`
		_, err := db.Db.Exec(insertStatement, item.UID, item.Module, item.Title, item.AsciiTitle)
		if err != nil {
		 	return nil, err
		}
	}
	return List()
}
