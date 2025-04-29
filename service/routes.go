package service

import (
	"awesomeProject2/types"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service types.PersonService
}

func NewHandler(service types.PersonService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes(router mux.Router) {
	//router.HandleFunc("/person", h.GetPersons).Methods(http.MethodGet)
	//router.HandleFunc("/person", h.NewPerson).Methods(http.MethodPost)
	//router.HandleFunc("/person", h.ChangePerson).Methods(http.MethodPut)
	//router.HandleFunc("/person", h.DeletePerson).Methods(http.MethodDelete)
}

func (h *Handler) GetPersons(w http.ResponseWriter, r http.Request) {

}

func (h *Handler) NewPerson(w http.ResponseWriter, r http.Request) {

}

func (h *Handler) ChangePerson(w http.ResponseWriter, r http.Request) {

}

func (h *Handler) DeletePerson(w http.ResponseWriter, r http.Request) {

}
