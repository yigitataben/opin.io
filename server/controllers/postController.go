package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yigitataben/opin.io/models"
	"github.com/yigitataben/opin.io/utils"
	"net/http"
	"strconv"
)

var _ models.Post

func GetPost(w http.ResponseWriter, r *http.Request) {
	newPosts := models.GetAllPosts()
	res, _ := json.Marshal(newPosts)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["postID"]
	ID, err := strconv.ParseInt(postID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing... (ERRCD: x01)")
	}
	postDetails, _ := models.GetPostByID(ID)
	res, _ := json.Marshal(postDetails)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	CreatePost := models.Post{}
	utils.ParseBody(r, CreatePost)
	p := CreatePost.CreatePost()
	res, _ := json.Marshal(p)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["postID"]
	ID, err := strconv.ParseInt(postID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing... (ERRCD: x02)")
	}
	post := models.DeletePost(ID)
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var updatePost = models.Post{}
	utils.ParseBody(r, updatePost)
	vars := mux.Vars(r)
	postID := vars["postID"]
	ID, err := strconv.ParseInt(postID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing... (ERRCD: x03)")
	}
	postDetails, db := models.GetPostByID(ID)
	if updatePost.Category != "" {
		postDetails.Category = updatePost.Category
	}
	if updatePost.Title != "" {
		postDetails.Title = updatePost.Title
	}
	//Burada hata çıkabilir (Post adından ötürü.)!!!
	if updatePost.Post != "" {
		postDetails.Post = updatePost.Post
	}
	db.Save(&postDetails)
	res, _ := json.Marshal(postDetails)
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
