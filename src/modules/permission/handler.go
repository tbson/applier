package permission

import (
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
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
	listRoute, err := List()
	if err != nil {
		fmt.Println(err)
	}

	output, err := json.Marshal(listRoute)
	if err != nil {
		fmt.Println(err)
	}

	linkRaw := "<%s>; rel=\"%s\""

	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}
	currentPath := protocol + "://" + r.Host + r.URL.Path
	link := make([]string, 0)
	link = append(link, fmt.Sprintf(linkRaw, currentPath, "next"))
	link = append(link, fmt.Sprintf(linkRaw, currentPath, "prev"))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Link", strings.Join(link, ", "))
	fmt.Fprintf(w, string(output))
}
