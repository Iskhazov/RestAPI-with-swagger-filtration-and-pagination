package service

import (
	"awesomeProject2/types"
	"awesomeProject2/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service types.PersonService
}

func NewHandler(service types.PersonService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes(router *mux.Router) {
	router.HandleFunc("/persons", h.GetPeople).Methods(http.MethodPost)
	router.HandleFunc("/person", h.CreatePerson).Methods(http.MethodPost)
	router.HandleFunc("/person/{id}", h.ChangePerson).Methods(http.MethodPut)
	router.HandleFunc("/person/{id}", h.DeletePerson).Methods(http.MethodDelete)
}

func (h *Handler) GetPeople(w http.ResponseWriter, r *http.Request) {
	var request types.GetPeopleRequest

	if err := utils.ParseJSON(r, &request); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var pageToken types.PageToken
	if request.PageToken != "" {
		decodedToken, err := utils.DecodeToken(request.PageToken)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		pageToken = *decodedToken
	} else {
		pageToken = types.PageToken{}
	}

	if request.Size > 500 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("page size must be less than 500"))
		return
	}

	fmt.Println(request.Filters)
	ps, err := h.service.GetPeople(pageToken, request.Size, request.Filters)
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
	params := mux.Vars(r)
	id := params["id"]

	personId, _ := strconv.Atoi(id)

	//check if post exists
	_, err := h.service.GetPersonById(personId)
	if err != nil {
		//Todo обработать
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	var ChangePerson types.DBPerson
	if err := utils.ParseJSON(r, &ChangePerson); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.PersonChange(personId, ChangePerson)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Success")
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	personId, _ := strconv.Atoi(id)

	//check if post exists
	person, err := h.service.GetPersonById(personId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	if err := utils.ParseJSON(r, &person); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.DeletePerson(personId)
	if err != nil {
		log.Fatal("Routes.go: Failed to delete person with id ", person.ID)
	}
	utils.WriteJSON(w, http.StatusOK, person.ID)

}
