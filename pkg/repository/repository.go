package repository

import "github.com/VMironiuk/offlinescore"

type Squad interface {
	getSquad() (offlinescore.Squad, error)
}

type Repository struct{
	Squad
}

func NewRepository() *Repository {
	return &Repository{}
}