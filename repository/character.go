package repository

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/supernova0730/rickandmortysync/model"
)

type CharacterRepository struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

func NewCharacterRepository(db *sql.DB) *CharacterRepository {
	return &CharacterRepository{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (repo *CharacterRepository) Insert(ch model.Character) (err error) {
	_, err = repo.builder.
		Insert("character").
		Columns("id", "name", "status", "species", "type", "gender", "image").
		Values(ch.ID, ch.Name, ch.Status, ch.Species, ch.Type, ch.Gender, ch.Image).
		RunWith(repo.db).Exec()
	return
}

func (repo *CharacterRepository) UpdateByID(id int64, ch model.Character) (err error) {
	_, err = repo.builder.
		Update("character").
		SetMap(map[string]interface{}{
			"name":    ch.Name,
			"status":  ch.Status,
			"species": ch.Species,
			"type":    ch.Type,
			"gender":  ch.Gender,
			"image":   ch.Image,
		}).
		Where(squirrel.Eq{"id": id}).
		RunWith(repo.db).
		Exec()
	return
}

func (repo *CharacterRepository) IsExistsByID(id int64) (result bool, err error) {
	count := 0
	err = repo.builder.
		Select("COUNT(1)").
		From("character").
		Where(squirrel.Eq{"id": id}).
		RunWith(repo.db).
		QueryRow().
		Scan(&count)
	result = count > 0
	return
}
