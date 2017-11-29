package permission

import (
	// "fmt"
	// "encoding/json"
	"util/tool"
)

type Permission struct {
	ID int "json:id"
	UID string "json:uid"
	Module string "json:module"
	Title string "json:title"
	AsciiTitle string "json:title"
	CreatedAt string "json:created_at"
	UpdatedAt string "json:updated_at"
}

func Generate() []tool.RawPermission {
	listRoute := tool.ParseRouter();
	return listRoute
}
