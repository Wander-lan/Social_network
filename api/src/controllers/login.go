package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/safety"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Authenticates a user on the API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userOnDatabase, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = safety.CheckPassword(userOnDatabase.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userOnDatabase.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))

}
