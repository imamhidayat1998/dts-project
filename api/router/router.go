package router

import (
	"dts-project/api/handler"
	"dts-project/pkg/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(
	serviceCenterHandler model.DataUseCase,
)  {
	r := mux.NewRouter()

	h := handler.NewServiceHandler(serviceCenterHandler)
	r.HandleFunc("/create",h.CreateUsersHandler).Methods("POST")



	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	err := http.ListenAndServe("",handlers.CORS(headers,methods,origins)(r))
	if err != nil{
		return
	}

}