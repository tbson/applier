package permission

import (
	"common/constant"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	// "github.com/gorilla/mux"
)

type API struct {
	Message string "json:message"
}

/*
func add(w http.ResponseWriter, r *http.Request) {
	permission := Permission{}
	permission.UID = urlParams["uid"]

	Add(permission)
}
*/

func sync(w http.ResponseWriter, r *http.Request) {
	/*
		urlParams := mux.Vars(r)
		name := urlParams["user"]
		helloMessage := "Permission " + name

		message := API{helloMessage}
		output, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		fmt.Fprintf(w, string(output))
	*/
	listRoute, err := Sync()
	if err != nil {
		fmt.Println(err)
	}

	output, err := json.Marshal(listRoute)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(output))
}

func list(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()

	startArr, ok := urlParams["start"]
	start := 0
	if ok {
		result, err := strconv.Atoi(startArr[0])
		if err != nil {
			start = 0
		} else {
			start = result
		}
	}

	directionArr, ok := urlParams["direction"]
	direction := ""
	if ok {
		direction = directionArr[0]
	}

	pOption := &constant.POption{start, direction}
	listRoute, err := List(pOption)
	if err != nil {
		fmt.Println(err)
	}

	firstId := 0
	lastId := 0

	if len(listRoute) == constant.PageSize {
		// Have enough results
		firstId = listRoute[0].ID
		lastId = listRoute[len(listRoute)-1].ID
	} else {
		// Not enough or missing result
		resultCount := len(listRoute)
		if direction == "next" {
			// origin = "right"
			if resultCount > 0 {
				firstId = listRoute[0].ID
			} else {
				firstId = pOption.Start
			}
		}
		if direction == "prev" {
			// origin = "left"
			if resultCount > 0 {
				lastId = listRoute[len(listRoute)-1].ID
			} else {
				lastId = pOption.Start
			}
		}
	}

	output, err := json.Marshal(listRoute)
	if err != nil {
		fmt.Println(err)
	}

	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}
	currentPath := protocol + "://" + r.Host + r.URL.Path
	link := make([]string, 0)
	if firstId != 0 {
		link = append(link, fmt.Sprintf(constant.HeaderLinkRaw, currentPath, firstId, "prev", "prev"))
	}
	if lastId != 0 {
		link = append(link, fmt.Sprintf(constant.HeaderLinkRaw, currentPath, lastId, "next", "next"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Link", strings.Join(link, ", "))

	fmt.Fprintf(w, string(output))
}
