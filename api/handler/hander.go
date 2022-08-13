package handler

import (
	"dts-project/api"
	"dts-project/pkg/model"
	"dts-project/pkg/model/request"
	"encoding/json"
	"fmt"
	"html/template"
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

func (s *ServiceHandler) Create(w http.ResponseWriter, r *http.Request) {

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

	api.Json(w, r, result, http.StatusCreated)

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

	//var requestData request.Artikel
	//decoder := json.NewDecoder(r.Body)
	//err := decoder.Decode(&requestData)
	//if err != nil {
	//	return
	//}
	//result, err := s.serviceUsecase.CreateArtikel(requestData)
	//if err != nil {
	//	return
	//}
	//api.Json(w, r, result, http.StatusOK)
	file := []string{
		"views/base.html",
		"views/create.html"}
	htmlTemplate, err := template.ParseFiles(file...)
	if err != nil {
		fmt.Print(err)
		return
	}
	err = htmlTemplate.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func (s *ServiceHandler) Index(w http.ResponseWriter, r *http.Request) {

	file := []string{
		"views/base.html",
		"views/index.html"}

	htmlTemplate, err := template.ParseFiles(file...)
	if err != nil {
		fmt.Print(err)
		return
	}
	var data []request.Artikel
	ResultData := map[string]interface{}{
		"notes": data,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", ResultData)
	if err != nil {
		fmt.Print(err)
		return
	}
}
