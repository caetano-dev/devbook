package controllers

import (
	"api/src/authentication"
	"api/src/base"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Login the user in the api
func Login(w http.ResponseWriter, r *http.Request) {
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
	db, error := base.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRespository(db)
	userSavedInDatabase, error := repository.FetchByEmail(user.Email)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.VerifyPassword(userSavedInDatabase.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	token, error := authentication.CreateToken(userSavedInDatabase.ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	w.Write([]byte(token))
}
