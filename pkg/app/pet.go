package app

import "github.com/Peter-Immanuel/fox-alpine/pkg/domain"

type petService struct {
	DB domain.PetDB
}

func NewPetService(db domain.PetDB) petService {
	return petService{
		DB: db,
	}
}

func (ps petService) Get(id int) (*domain.Pet, error) {
	return ps.DB.Get(id)
}

func (ps petService) List(query ...string) ([]*domain.Pet, error) {
	return ps.DB.List(query...)
}

func (ps petService) Create(pet *domain.Pet) error {
	return ps.DB.Create(pet)
}

func (ps petService) Delete(id any) error {
	return ps.DB.Delete(id)
}
