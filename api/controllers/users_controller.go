package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/obasajujoshua31/blogos/api/database"
	"github.com/obasajujoshua31/blogos/api/models"
	"github.com/obasajujoshua31/blogos/api/repository"
	"github.com/obasajujoshua31/blogos/api/repository/crud"
	"github.com/obasajujoshua31/blogos/api/responses"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		users, err = usersRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, users)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	convID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err := usersRepository.FindByID(uint32(convID))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, user)
	}(repo)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err = usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	convID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(body, &user)

	if user.Password != "" {
		responses.ERROR(w, http.StatusForbidden, errors.New("You cannot change your password"))
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err = usersRepository.Update(uint32(convID), user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, user)
	}(repo)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
