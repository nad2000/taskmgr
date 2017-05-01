package main

import (
	"log"
	"net/http"

	"taskmgr/common"
	"taskmgr/routers"

	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	r := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Listening...")
	http.ListenAndServe(":8080", n)
}
