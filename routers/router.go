package routers

import (
	"github.com/gorilla/mux"
)

// InitRouters initilizes routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetTaskRouters(router)
	return router
}
