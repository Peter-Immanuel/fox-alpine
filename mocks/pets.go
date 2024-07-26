package mocks

import "github.com/Peter-Immanuel/fox-alpine/pkg/domain"

type PetService struct {
	GetPetResp   domain.Pet
	GetPetErr    error
	ListPetResp  []*domain.Pet
	ListPetErr   error
	CreatePetErr error
	DeletePetErr error
}

func (ps PetService) Get(domain.PetID) (*domain.Pet, error) {
	return &ps.GetPetResp, ps.GetPetErr
}

func (ps PetService) List(string) ([]*domain.Pet, error) {
	return ps.ListPetResp, ps.ListPetErr
}

func (ps PetService) Create(*domain.Pet) error {
	return ps.CreatePetErr
}

func (ps PetService) Delete(domain.PetID) error {
	return ps.DeletePetErr
}
