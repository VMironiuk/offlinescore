package repository

import "github.com/VMironiuk/offlinescore"

type SquadRepository struct{}

func NewSquadRepository() *SquadRepository {
	return &SquadRepository{}
}

func (r *SquadRepository) GetSquad() (offlinescore.Squad, error) {
	squad := offlinescore.Squad{
		Id: "4009",
	}

	return squad, nil
}
