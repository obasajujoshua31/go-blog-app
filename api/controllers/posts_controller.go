package controllers

import (
	"encoding/json"
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	err = post.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		post, err = postsRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := []models.Post{}
	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		posts, err = postsRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, posts)
	}(repo)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
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

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		user, err := postsRepository.FindByID(uint32(convID))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, user)
	}(repo)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var post models.Post
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
	err = json.Unmarshal(body, &post)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postsRepository repository.PostRepository) {
		post, err = postsRepository.Update(uint32(convID), post)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, post)
	}(repo)
}
