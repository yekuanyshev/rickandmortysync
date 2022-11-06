package processor

import (
	"github.com/supernova0730/rickandmortysync/client"
	"github.com/supernova0730/rickandmortysync/model"
	"github.com/supernova0730/rickandmortysync/repository"

	"golang.org/x/sync/errgroup"
)

type Processor struct {
	client *client.Client
	repo   *repository.CharacterRepository
}

func New(client *client.Client, repo *repository.CharacterRepository) *Processor {
	return &Processor{
		client: client,
		repo:   repo,
	}
}

func (p *Processor) Sync() error {
	info, err := p.client.GetInfo()
	if err != nil {
		return err
	}

	var g errgroup.Group
	for page := 1; page < info.Pages; page++ {
		page := page
		g.Go(func() error {
			return p.synchronise(page)
		})
	}
	return g.Wait()
}

func (p *Processor) SyncByNumWorkers(numWorkers int) error {
	info, err := p.client.GetInfo()
	if err != nil {
		return err
	}

	size := p.getSizeByPagesAndNumOfWorkers(info.Pages, numWorkers)

	var g errgroup.Group
	for i := 0; i < numWorkers; i++ {
		i := i
		g.Go(func() error {
			start, end := p.pagination(i, size, info.Pages)
			if start > end {
				return nil
			}

			return p.synchroniseByPages(start, end)
		})
	}
	return g.Wait()
}

func (p *Processor) synchroniseByPages(startPage, endPage int) error {
	for page := startPage; page <= endPage; page++ {
		err := p.synchronise(page)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Processor) synchronise(page int) error {
	characters, err := p.client.GetCharactersByPage(page)
	if err != nil {
		return err
	}

	for _, character := range characters {
		exists, err := p.repo.IsExistsByID(character.ID)
		if err != nil {
			return err
		}

		c := p.toModel(character)
		if exists {
			err = p.repo.UpdateByID(character.ID, c)
		} else {
			err = p.repo.Insert(c)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Processor) pagination(i, size, max int) (start, end int) {
	start = i*size + 1
	end = (i + 1) * size
	if end > max {
		end = max
	}
	return
}

func (p *Processor) getSizeByPagesAndNumOfWorkers(pages int, numOfWorkers int) int {
	size := pages / numOfWorkers
	if pages%numOfWorkers != 0 {
		size += 1
	}
	return size
}

func (p *Processor) toModel(character client.Character) model.Character {
	return model.Character{
		ID:      character.ID,
		Name:    character.Name,
		Status:  character.Status,
		Species: character.Species,
		Type:    character.Type,
		Gender:  character.Gender,
		Image:   character.Gender,
	}
}
