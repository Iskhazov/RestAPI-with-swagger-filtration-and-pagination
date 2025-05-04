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

// Handler struct holds the reference to the service layer
type Handler struct {
	service types.PersonService
}

// NewHandler creates a new Handler instance
func NewHandler(service types.PersonService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes(router *mux.Router) {
	router.HandleFunc("/persons", h.GetPeople).Methods(http.MethodPost)
	router.HandleFunc("/person", h.CreatePerson).Methods(http.MethodPost)
	router.HandleFunc("/person/{id}", h.ChangePerson).Methods(http.MethodPut)
	router.HandleFunc("/person/{id}", h.DeletePerson).Methods(http.MethodDelete)
}

// GetPeople godoc
// @Summary Get all people
// @Description The request accepts filters, a pagination token, and the number of items to return (size). The size must not exceed 100.
// @Description For the first request, send an empty `page_token`. The response will include a `next_page_token`, which should be used in the subsequent request to fetch the next page.
// @Description You can filter results by age and gender.
// @Description To filter by age, set `"age"` as the field and provide one or two values: e.g., `["20"]` for a single value or `["20", "50"]` for a range.
// @Description To filter by gender, set `"gender"` as the field and provide one value: `"male"` or `"female"`.
// @Description To retrieve both genders, omit the gender filter entirely.
// @Tags people
// @Accept  json
// @Produce  json
// @Param request body types.GetPeopleRequest true "Pagination and filtering options"
// @Success 200 {array} types.GetPeopleResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /persons [post]
func (h *Handler) GetPeople(w http.ResponseWriter, r *http.Request) {

	var request types.GetPeopleRequest

	if err := utils.ParseJSON(r, &request); err != nil {
		log.Printf("[ERROR] Failed to parse request JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var pageToken types.PageToken
	if request.PageToken != "" {
		decodedToken, err := utils.DecodeToken(request.PageToken)
		if err != nil {
			log.Printf("[ERROR] Failed to decode page token: %v", err)
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		pageToken = *decodedToken
	} else {
		pageToken = types.PageToken{}
	}

	if request.Size > 100 {
		log.Printf("[WARN] Page size exceeded: %d", request.Size)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("page size must be less than 100"))
		return
	}

	ps, err := h.service.GetPeople(pageToken, request.Size, request.Filters)
	if err != nil {
		log.Printf("[ERROR] Failed to get people: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("[INFO] Successfully retrieved %d people", len(ps.People))
	utils.WriteJSON(w, http.StatusOK, ps)
}

// CreatePerson godoc
// @Summary Create a new person
// @Description Creates a new person, enriches data using external APIs, and saves the result to the database.
// @Tags people
// @Accept  json
// @Produce  json
// @Param person body types.NewPerson true "Person to create"
// @Success 201 {string} string "Success"
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /person [post]
func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {

	var person types.NewPerson
	if err := utils.ParseJSON(r, &person); err != nil {
		log.Printf("[ERROR] Failed to parse person JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.service.CreatePerson(person); err != nil {
		log.Printf("[ERROR] Failed to create person: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("[INFO] Person created successfully")
	utils.WriteJSON(w, http.StatusCreated, "Success")
}

// ChangePerson godoc
// @Summary Update a person
// @Description Updates the information of a person by ID
// @Tags people
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Param person body types.SwagPerson true "Updated person data"
// @Success 200 {string} string "Success"
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /person/{id} [put]
func (h *Handler) ChangePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	personId, _ := strconv.Atoi(id)

	// Check if the person exists before updating
	_, err := h.service.GetPersonById(personId)
	if err != nil {
		log.Printf("[WARN] Person with ID %d not found", personId)
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	var updatedPerson types.DBPerson
	if err := utils.ParseJSON(r, &updatedPerson); err != nil {
		log.Printf("[ERROR] Failed to parse updated person JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.service.PersonChange(personId, updatedPerson); err != nil {
		log.Printf("[ERROR] Failed to update person with ID %d: %v", personId, err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("[INFO] Person with ID %d updated successfully", personId)
	utils.WriteJSON(w, http.StatusOK, "Success")
}

// DeletePerson godoc
// @Summary Delete a person
// @Description Deletes a person by ID
// @Tags people
// @Accept  json
// @Produce  json
// @Param id path int true "Person ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /person/{id} [delete]
func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	personId, _ := strconv.Atoi(id)

	// Check if person exists before deletion
	person, err := h.service.GetPersonById(personId)
	if err != nil {
		log.Printf("[WARN] Person with ID %d not found", personId)
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	if err := h.service.DeletePerson(person.ID); err != nil {
		log.Printf("[ERROR] Failed to delete person with ID %d: %v", personId, err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("[INFO] Person with ID %d deleted successfully", personId)
	utils.WriteJSON(w, http.StatusOK, "Success")
}
