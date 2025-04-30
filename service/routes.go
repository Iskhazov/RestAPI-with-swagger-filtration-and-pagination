package service

import (
	"awesomeProject2/types"
	"awesomeProject2/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	service types.PersonService
}

func NewHandler(service types.PersonService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes(router *mux.Router) {
	router.HandleFunc("/person", h.GetPeople).Methods(http.MethodGet)
	router.HandleFunc("/person", h.CreatePerson).Methods(http.MethodPost)
	router.HandleFunc("/person", h.ChangePerson).Methods(http.MethodPut)
	router.HandleFunc("/person", h.DeletePerson).Methods(http.MethodDelete)
}

func (h *Handler) GetPeople(w http.ResponseWriter, r *http.Request) {
	ps, err := h.service.GetPeople()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person types.NewPerson

	if err := utils.ParseJSON(r, &person); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err := h.service.CreatePerson(person)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, "Success")
}

func (h *Handler) ChangePerson(w http.ResponseWriter, r *http.Request) {
	var person types.DBPerson

	if err := utils.ParseJSON(r, &person); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.service.PersonChange(person)
	if err != nil {
		log.Fatal("Routes.go: Failed to change person with id ", person.ID)
	}
	utils.WriteJSON(w, http.StatusOK, "Success")
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	var person types.Person

	if err := utils.ParseJSON(r, &person); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.service.DeletePerson(person.ID)
	if err != nil {
		log.Fatal("Routes.go: Failed to delete person with id ", person.ID)
	}
	utils.WriteJSON(w, http.StatusOK, person.ID)

}
