package work

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllWorks() ([]Work, error) {
	var works []Work
	if err := r.db.Preload("Gallery").Find(&works).Error; err != nil {
		return nil, err
	}

	return works, nil
}
func (r *Repository) GetWork(id uint) (*Work, error) {
	var work Work
	if err := r.db.Preload("Gallery").First(&work, id).Error; err != nil {
		return nil, err
	}
	return &work, nil
}
