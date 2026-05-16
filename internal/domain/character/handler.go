package character

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) RegisterRoutes(router chi.Router) {
	router.Route("/characters", func(r chi.Router) {
		r.Get("/", handler.index)
		r.Post("/", handler.store)
		r.Get("/{id}", handler.show)
		r.Put("/{id}", handler.update)
		r.Delete("/{id}", handler.destroy)
	})
}

func (handler *Handler) index(responseWriter http.ResponseWriter, request *http.Request) {
	characters, err := handler.service.GetAll()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(characters)
}

func (handler *Handler) show(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		http.Error(responseWriter, "invalid id", http.StatusBadRequest)
		return
	}

	character, err := handler.service.GetByID(uint(id))
	if err != nil {
		http.Error(responseWriter, "character not found", http.StatusNotFound)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(character)
}

func (handler *Handler) store(responseWriter http.ResponseWriter, request *http.Request) {
	var character Character
	if err := json.NewDecoder(request.Body).Decode(&character); err != nil {
		http.Error(responseWriter, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := handler.service.Create(&character); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(character)
}

func (handler *Handler) update(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		http.Error(responseWriter, "invalid id", http.StatusBadRequest)
		return
	}

	existing, err := handler.service.GetByID(uint(id))
	if err != nil {
		http.Error(responseWriter, "character not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(request.Body).Decode(existing); err != nil {
		http.Error(responseWriter, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := handler.service.Update(existing); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(existing)
}

func (handler *Handler) destroy(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		http.Error(responseWriter, "invalid id", http.StatusBadRequest)
		return
	}

	if err := handler.service.Delete(uint(id)); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
