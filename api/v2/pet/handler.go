package pet

import (
	"encoding/json"
	"net/http"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	pets, _ := h.service.GetAllPets()
	response, _ := json.Marshal(pets)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h handler) GetById(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}
