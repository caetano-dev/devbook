package controllers

import (
	"api/src/authentication"
	"api/src/base"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}
	var post models.Post
	if error = json.Unmarshal(requestBody, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	post.AuthorID = userID
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewPostRepository(db)
	post.ID, error = repository.Create(post)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// FetchPosts fetches all posts
func FetchPosts(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewPostRepository(db)
	posts, error := repository.Fetch(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)

}

// FetchPost fetches a single post
func FetchPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewPostRepository(db)
	post, error := repository.FetchByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, post)

}

// UpdatePost updates a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	postSavedInDatabase, error := repository.FetchByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if postSavedInDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("You can't update a post that is not yours"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	if error = json.Unmarshal(requestBody, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.Update(postID, post); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	postSavedInDatabase, error := repository.FetchByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	if postSavedInDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("You can't delete a post that is not yours"))
		return
	}
	if error = repository.Delete(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}

// FetchPostByUser fetches all posts by a user
func FetchPostByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewPostRepository(db)
	posts, error := repository.FetchPostByUser(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, posts)

}

// LikePost likes a post
func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewPostRepository(db)
	if error = repository.Like(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// DislikePost removes the like from a post
func DislikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewPostRepository(db)
	if error = repository.Dislike(postID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
