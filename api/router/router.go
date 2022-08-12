package router

import (
	"dts-project/api/handler"
	"dts-project/pkg/model"
	"dts-project/pkg/model/request"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)


func Router(url request.URL, serviceCenterHandler model.DataUseCase, ) {
	RouterUser := mux.NewRouter()

	h := handler.NewServiceHandler(serviceCenterHandler)
	RouterUser.HandleFunc("/create",h.CreateUsersHandler).Methods("POST")
	RouterUser.HandleFunc("/update",h.UpdateUsersHandlers).Methods("PUT")



	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	fmt.Printf("Server listening %s:%s",url.Host,url.Port)
	err := http.ListenAndServe(url.Host+":"+url.Port,handlers.CORS(headers,methods,origins)(RouterUser))
	if err != nil{
		return
	}

}