package config

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"util/tool"
)

type API struct {
	Message string "json:message"
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(tool.ParseRouter())
	urlParams := mux.Vars(r)
	name := urlParams["user"]
	helloMessage := "Config " + name

	message := API{helloMessage}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output))
}
