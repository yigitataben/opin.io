package routes

import (
	"github.com/gorilla/mux"
	"github.com/yigitataben/opin.io/controllers"
	//"github.com/yigitataben/opin.io/server/controllers"
)

var RegisterPostRoutes = func(router mux.Router) {
	router.HandleFunc("/post/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post/", controllers.GetPost).Methods("GET")
	router.HandleFunc("/post/{postID}", controllers.GetPostByID).Methods("GET")
	router.HandleFunc("/post/{postID}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{postID}", controllers.DeletePost).Methods("DELETE")
}
