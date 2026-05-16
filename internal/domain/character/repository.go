package character

import "gorm.io/gorm"

type Repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) *Repository {
	return &Repository{database: database}
}

func (repository *Repository) FindAll() ([]Character, error) {
	var characters []Character
	result := repository.database.Find(&characters)
	return characters, result.Error
}

func (repository *Repository) FindByID(id uint) (*Character, error) {
	var character Character
	result := repository.database.First(&character, id)
	return &character, result.Error
}

func (repository *Repository) Create(character *Character) error {
	return repository.database.Create(character).Error
}

func (repository *Repository) Update(character *Character) error {
	return repository.database.Save(character).Error
}

func (repository *Repository) Delete(id uint) error {
	return repository.database.Delete(&Character{}, id).Error
}
