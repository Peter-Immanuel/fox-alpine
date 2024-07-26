package app

import "github.com/Peter-Immanuel/fox-alpine/pkg/domain"

type petService struct {
	DB domain.PetDB
}

func NewPetService(db domain.PetDB) domain.PetService {
	return petService{
		DB: db,
	}
}

func (ps petService) Get(id domain.PetID) (*domain.Pet, error) {
	return ps.DB.Get(id)
}

func (ps petService) List(category string) ([]*domain.Pet, error) {
	return ps.DB.List(category)
}

func (ps petService) Create(pet *domain.Pet) (*domain.Pet, error) {
	return ps.DB.Create(pet)
}

func (ps petService) Delete(id domain.PetID) error {
	return ps.DB.Delete(id)
}
