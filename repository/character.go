package repository

import (
	"github.com/supernova0730/rickandmortysync/model"
	"gorm.io/gorm"
)

type CharacterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) *CharacterRepository {
	return &CharacterRepository{db: db}
}

func (repo *CharacterRepository) ByID(id int64) (result model.Character, err error) {
	err = repo.db.Model(&model.Character{}).
		Where("id = ?", id).
		First(&result).Error
	return
}

func (repo *CharacterRepository) Insert(character model.Character) (err error) {
	err = repo.db.Model(&model.Character{}).
		Create(&character).Error
	return
}

func (repo *CharacterRepository) UpdateByID(id int64, character model.Character) (err error) {
	err = repo.db.Model(&model.Character{}).
		Where("id = ?", id).
		Updates(&character).Error
	return
}

func (repo *CharacterRepository) IsExistsByID(id int64) (result bool, err error) {
	err = repo.db.Model(&model.Character{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&result).Error
	return
}
