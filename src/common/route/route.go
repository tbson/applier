package route

import (
	"github.com/gorilla/mux"
)


var MainRouter = mux.NewRouter()
var Router = MainRouter.PathPrefix("/api/").Subrouter()
