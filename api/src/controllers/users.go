package controllers

import (
	"api/src/authentication"
	"api/src/base"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("register"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRespository(db)
	user.ID, error = repository.Create(user)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

//FetchUsers fetches users
func FetchUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRespository(db)
	users, error := repository.Fetch(nameOrNick)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

//FetchUser fetches a user
func FetchUser(w http.ResponseWriter, r *http.Request) {
	paramenters := mux.Vars(r)
	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)

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

	repository := repositories.NewUserRespository(db)
	user, error := repository.FetchByID(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

//UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	paramenters := mux.Vars(r)
	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	userIDInToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if userID != userIDInToken {
		responses.Error(w, http.StatusForbidden, errors.New("It is not possible to update a user that is not yours"))
		return
	}
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}
	var user models.User

	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("edition"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRespository(db)
	if error = repository.Update(userID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	paramenters := mux.Vars(r)

	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userIDInToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	if userID != userIDInToken {
		responses.Error(w, http.StatusForbidden, errors.New("You can't delete an user that is not yours"))
		return

	}

	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRespository(db)

	if error = repository.Delete(userID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}

//FollowUser lets an user follow another
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserID(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	paramenters := mux.Vars(r)
	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Impossible to follow yourself"))
		return
	}
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRespository(db)
	if error = repository.Follow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

//unfollowUser lets an user unfollow another
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	paramenters := mux.Vars(r)
	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Impossible to unfollow yourself"))
		return
	}

	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRespository(db)
	if error = repository.Unfollow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

//FetchFollowers shows all followers from user
func FetchFollowers(w http.ResponseWriter, r *http.Request) {
	paramenters := mux.Vars(r)
	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)

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

	repository := repositories.NewUserRespository(db)
	followers, error := repository.FetchFollowers(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

//FetchFollowing shows all accounts the user follows
func FetchFollowing(w http.ResponseWriter, r *http.Request) {
	paramenters := mux.Vars(r)
	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)

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

	repository := repositories.NewUserRespository(db)
	users, error := repository.FetchFollowing(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

//UpdatePassword from user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDInToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	paramenters := mux.Vars(r)
	userID, error := strconv.ParseUint(paramenters["userID"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	if userIDInToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("It is not possible to update a password form another user"))
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	var password models.Password
	if error = json.Unmarshal(requestBody, &password); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRespository(db)
	PasswordSavedInDatabase, error := repository.FetchPassword(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.VerifyPassword(PasswordSavedInDatabase, password.CurrentPassword); error != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("The current password is wrong"))
		return
	}

	passwordWithHash, error := security.Hash(password.NewPassword)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.UpdatePassword(userID, string(passwordWithHash)); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
