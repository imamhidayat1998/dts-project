package handler

import (
	"dts-project/api"
	"dts-project/pkg/model"
	"dts-project/pkg/model/request"
	"encoding/json"
	"net/http"
)

type ServiceHandler struct {
	serviceUsecase model.DataUseCase
}

func NewServiceHandler(
	serviceCenterUsecase model.DataUseCase,
) *ServiceHandler {

	return &ServiceHandler{
		serviceUsecase: serviceCenterUsecase,
	}
}

func (s *ServiceHandler) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {

	var requestData request.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		return
	}
	result, err := s.serviceUsecase.CreateUsers(requestData)
	if err != nil {
		return
	}
	api.Json(w, r, result, http.StatusOK)

}
func (s *ServiceHandler) UpdateUsersHandlers(w http.ResponseWriter, r *http.Request) {

	var requestData request.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		return
	}
	err = s.serviceUsecase.UpdateUsers(requestData)
	if err != nil {
		return
	}
	api.Json(w, r, "Updated Success", http.StatusOK)
}
func (s *ServiceHandler) CreateArtikelHandler(w http.ResponseWriter, r *http.Request) {

	var requestData request.Artikel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		return
	}
	result, err := s.serviceUsecase.CreateArtikel(requestData)
	if err != nil {
		return
	}
	api.Json(w, r, result, http.StatusOK)

}
