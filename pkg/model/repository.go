package model

import "dts-project/pkg/model/request"

type DataRepository interface {
	CreateUser(params request.User) (string, error)
	UpdateUser(params request.User) error
	SelectId(params request.User) (*request.User, error)
	CreateArtikel(params request.Artikel) (string, error)
	GetArtikel() ([]request.Artikel, error)
}
