package service

import (
	"github.com/VMironiuk/offlinescore"
	"github.com/VMironiuk/offlinescore/pkg/repository"
)

type Squad interface {
	GetSquad() (offlinescore.Squad, error)
}

type Service struct {
	Squad
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Squad: NewSquadService(repo.Squad),
	}
}
