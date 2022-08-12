package model

import "dts-project/pkg/model/request"

type DataUseCase interface {
	UpdateUsers(params request.User) error
	CreateUsers(params request.User) (string, error)
	CreateArtikel(params request.Artikel) (string,error)
}
