package crud

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/obasajujoshua31/blogos/api/models"
	"github.com/obasajujoshua31/blogos/api/utils/channels"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error
	var users []models.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Find(&users).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return users, nil
	}
	return []models.User{}, err
}

func (r *repositoryUsersCRUD) FindByID(id uint32) (models.User, error) {
	var err error
	var user models.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Where("ID = ?", id).Take(&user).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("User not found")
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindByEmail(email string) (models.User, error) {
	var err error
	var user models.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Where("email = ?", email).Take(&user).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("User not found")
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) Update(id uint32, user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Updates(user).Error

		if err != nil {
			ch <- false
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("User not found")
	}
	return models.User{}, err
}
