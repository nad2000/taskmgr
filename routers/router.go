package routers

import (
	"github.com/gorilla/mux"
)

// InitRouters initilizes routes
func InitRoutes() *mux.Router {
	return mux.NewRouter().StrictSlash(false)
}
