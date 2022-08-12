package model

import "dts-project/pkg/model/request"
type DataRepository interface {
	CreateUser(params request.User) (string, error)
	UpdateUser(params request.User) error

}

