package crud

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/obasajujoshua31/blogos/api/models"
	"github.com/obasajujoshua31/blogos/api/utils/channels"
)

type repositoryPostsCRUD struct {
	db *gorm.DB
}

func NewRepositoryPostsCRUD(db *gorm.DB) *repositoryPostsCRUD {
	return &repositoryPostsCRUD{db}
}

func (r *repositoryPostsCRUD) Save(post models.Post) (models.Post, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Post{}).Create(&post).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return post, nil
	}
	return models.Post{}, err
}

func (r *repositoryPostsCRUD) FindAll() ([]models.Post, error) {
	var err error
	var posts []models.Post
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Post{}).Find(&posts).Error

		if err != nil {
			ch <- false
		}

		if len(posts) > 0 {

			for i, _ := range posts {
				err = r.db.Debug().Model(&models.User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
				if err != nil {
					ch <- false
				}
			}

		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return posts, nil
	}
	return []models.Post{}, err
}

func (r *repositoryPostsCRUD) FindByID(id uint32) (models.Post, error) {
	var err error
	var post models.Post
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Post{}).Where("ID = ?", id).Take(&post).Error

		if err != nil {
			ch <- false
		}

		err = r.db.Debug().Model(&models.User{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
		if err != nil {
			ch <- false
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return post, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Post{}, errors.New("Post not found")
	}
	return models.Post{}, err
}

func (r *repositoryPostsCRUD) Update(id uint32, post models.Post) (models.Post, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Post{}).Updates(post).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return post, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Post{}, errors.New("Post not found")
	}
	return models.Post{}, err
}
