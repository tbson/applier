package permission

import (
	"fmt"
	// "encoding/json"
	"common/constant"
	"common/db"
	"common/route"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"strings"
	"util/tool"
)

type RawPermission struct {
	UID        string "json:uid"
	Module     string "json:module"
	Title      string "json:title"
	AsciiTitle string "json:title"
}

type Permission struct {
	ID         int    "json:id"
	UID        string "json:uid"
	Module     string "json:module"
	Title      string "json:title"
	AsciiTitle string "json:title"
	CreatedAt  string "json:created_at"
	UpdatedAt  string "json:updated_at"
}

func ParseRouter() []*RawPermission {
	var parsedRouter = make([]*RawPermission, 0)
	route.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeUid := route.GetName()
		if routeUid != "" {
			rawPermission := &RawPermission{}
			routeUidArr := strings.Split(routeUid, "_")

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

func List(pOption *constant.POption) ([]*Permission, error) {
	var result = make([]*Permission, 0)
	var rows *sql.Rows
	var err error
	order := "ASC"
	if pOption.Direction == "prev" {
		order = "DESC"
	}

	firstQuery := `
		SELECT id, uid, module, title, ascii_title, created_at, updated_at
		FROM permission
		ORDER BY id %s
		LIMIT %d
	`
	firstQuery = fmt.Sprintf(firstQuery, order, constant.PageSize)

	pageQuery := `
		SELECT id, uid, module, title, ascii_title, created_at, updated_at
		FROM permission
		WHERE id %s $1
		ORDER BY id %s
		LIMIT %d
	`
	pageQuery = fmt.Sprintf(pageQuery, tool.DirectionParse(pOption.Direction), order, constant.PageSize)

	if pOption.Start == 0 && pOption.Direction == "" {
		rows, err = db.Db.Query(firstQuery)
	} else {
		rows, err = db.Db.Query(pageQuery, pOption.Start)
	}

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
		if pOption.Direction == "next" {
			result = append(result, &item)
		} else {
			newResult := make([]*Permission, len(result)+1)
			newResult[0] = &item
			copy(newResult[1:], result)
			result = newResult
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func Sync() ([]*Permission, error) {
	listRoute := ParseRouter()
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
	pOption := &constant.POption{3, "next"}
	return List(pOption)
}
