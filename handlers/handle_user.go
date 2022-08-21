package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/denizedizcan/REST-API-User-Management-Service/models"
	"github.com/denizedizcan/REST-API-User-Management-Service/responses"
	"github.com/gorilla/mux"
)

// create user handler
func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate("Create"); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()

	err = user.SaveUser(h.DB)
	if errors.Is(err, gorm.ErrInvalidData) {
		responses.ERROR(w, http.StatusForbidden, errors.New("User with that email already exists"))
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, &user)
}

// show user details handler
func (h handler) ShowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.UserID = id

	err = user.FindUser(h.DB)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		responses.ERROR(w, http.StatusNotFound, errors.New("User with that id does not exist"))
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, &user)
}

//find all users
func (h handler) ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.FindAllUsers(h.DB)

	if len(users) == 0 {
		responses.ERROR(w, http.StatusNotFound, errors.New("User with that id does not exist"))
		return
	}

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.UserID = id
	err = user.DeleteUser(h.DB)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		responses.ERROR(w, http.StatusNotFound, errors.New("User with that id does not exist"))
		return
	}

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "")
}

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	var user models.User
	user.UserID = id

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	var m map[string]interface{}
	json.Unmarshal(body, &m)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	fmt.Println(m)
	if err := user.UpdateUser(m, h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err := user.FindUser(h.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, &user)
}
