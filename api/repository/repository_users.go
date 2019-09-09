package repository

import "github.com/obasajujoshua31/blogos/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindByID(uint32) (models.User, error)
	Update(uint32, models.User) (models.User, error)
	FindByEmail(string) (models.User, error)
	// Delete(uint32) (int64, error)
}
