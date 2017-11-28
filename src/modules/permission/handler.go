package permission

import (
	"fmt"
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

func index(w http.ResponseWriter, r *http.Request) {
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
	listRoute := Generate()
	output, err := json.Marshal(listRoute)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output))
}
