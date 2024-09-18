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

func (r *Repository) CreateWork(work *Work) error {
	return r.db.Create(work).Error
}
func (r *Repository) UpdateWork(work *Work) error {
	return r.db.Save(work).Error
}
func (r *Repository) DeleteWork(id uint) error {
	if err := r.db.Where("work_id = ?", id).Delete(&Gallery{}).Error; err != nil {
		return err
	}
	return r.db.Delete(&Work{}, id).Error
}
func (r *Repository) CreateGallery(gallery *Gallery) error {
	return r.db.Create(gallery).Error
}
