package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yigitataben/opin.io/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterPostRoutes(*r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}
