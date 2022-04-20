package repository

import "github.com/VMironiuk/offlinescore"

type SquadRepository struct{}

func NewSquadRepository() *SquadRepository {
	return &SquadRepository{}
}

func (r *SquadRepository) GetSquad() (offlinescore.Squad, error) {
	squad := makeSquad()
	return squad, nil
}
