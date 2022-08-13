package router

import (
	"dts-project/api/handler"
	"dts-project/pkg/model"
	"dts-project/pkg/model/request"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router(url request.URL, serviceCenterHandler model.DataUseCase) {
	RouterUser := mux.NewRouter()

	h := handler.NewServiceHandler(serviceCenterHandler)
	RouterUser.HandleFunc("/create", h.Create).Methods("POST")
	RouterUser.HandleFunc("/update", h.UpdateUsersHandlers).Methods("PUT")
	RouterUser.HandleFunc("/create-artikel", h.CreateArtikelHandler).Methods("POST")
	RouterUser.HandleFunc("/", h.Index).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	fmt.Printf("Server listening %s:%s", url.Host, url.Port)
	err := http.ListenAndServe(url.Host+":"+url.Port, handlers.CORS(headers, methods, origins)(RouterUser))
	if err != nil {
		return
	}

}
