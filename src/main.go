package main

import (
	"net/http"
	"common/route"
    _"common/db"
	_"modules/config"
	_"modules/permission"
)

func main() {
	http.Handle("/", route.Router)
	http.ListenAndServe(":8000", nil)
}
