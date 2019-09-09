package repository

import "github.com/obasajujoshua31/blogos/api/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindByID(uint32) (models.Post, error)
	Update(uint32, models.Post) (models.Post, error)
	// Delete(uint32) (int64, error)
}
