package service

import "github.com/VMironiuk/offlinescore"

type Squad interface {
	getSquad() (offlinescore.Squad, error)
}

type Service struct {
	Squad
}

func NewService() *Service {
	return &Service{}
}
