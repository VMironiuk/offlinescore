package service

import (
	"github.com/VMironiuk/offlinescore"
	"github.com/VMironiuk/offlinescore/pkg/repository"
)

type SquadService struct {
	repo repository.Squad
}

func NewSquadService(repo repository.Squad) *SquadService {
	return &SquadService{
		repo: repo,
	}
}

func (s *SquadService) GetSquad() (offlinescore.Squad, error) {
	return s.repo.GetSquad()
}
