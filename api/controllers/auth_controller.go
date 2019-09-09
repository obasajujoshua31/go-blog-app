package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/obasajujoshua31/blogos/api/auth"
	"github.com/obasajujoshua31/blogos/api/database"
	"github.com/obasajujoshua31/blogos/api/models"
	"github.com/obasajujoshua31/blogos/api/repository"
	"github.com/obasajujoshua31/blogos/api/repository/crud"
	"github.com/obasajujoshua31/blogos/api/responses"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
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

	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userPassword := user.Password

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err = usersRepository.FindByEmail(user.Email)

		if err != nil {
			responses.ERROR(w, http.StatusForbidden, errors.New("Invalid Email or Password"))
			return
		}

		err = user.ConfirmPassword(userPassword)

		if err != nil {
			responses.ERROR(w, http.StatusForbidden, errors.New("Invalid Email or Password"))
			return
		}

		token, err := auth.GenerateToken(user)

		if err != nil {
			responses.ERROR(w, http.StatusServiceUnavailable, errors.New("Unavailable to Login"))
		}

		response := map[string]string{
			"token":   token,
			"message": "You are successfully Logged In",
		}

		responses.JSON(w, http.StatusCreated, response)
	}(repo)

}
