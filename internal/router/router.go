package router

import (
	"encoding/json"
	"net/http"

	"dnd-sheet-api/internal/domain/character"
	"dnd-sheet-api/internal/middleware"
	"dnd-sheet-api/internal/proxy"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func New(database *gorm.DB, dnd5eClient *proxy.Dnd5eClient) *chi.Mux {
	router := chi.NewRouter()

	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	router.Use(middleware.Cors)

	router.Get("/health", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"status":"ok"}`))
	})

	// Character routes
	characterRepository := character.NewRepository(database)
	characterService := character.NewService(characterRepository)
	characterHandler := character.NewHandler(characterService)
	characterHandler.RegisterRoutes(router)

	// 5e-bits proxy routes
	router.Route("/dnd5e", func(r chi.Router) {
		r.Get("/races", func(responseWriter http.ResponseWriter, request *http.Request) {
			races, err := dnd5eClient.GetRaces()
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(races)
		})

		r.Get("/races/{index}", func(responseWriter http.ResponseWriter, request *http.Request) {
			raceData, err := dnd5eClient.GetRaceData(chi.URLParam(request, "index"))
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(raceData)
		})

		r.Get("/classes", func(responseWriter http.ResponseWriter, request *http.Request) {
			classes, err := dnd5eClient.GetClasses()
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(classes)
		})

		r.Get("/classes/{index}", func(responseWriter http.ResponseWriter, request *http.Request) {
			classData, err := dnd5eClient.GetClassData(chi.URLParam(request, "index"))
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(classData)
		})

		r.Get("/spells", func(responseWriter http.ResponseWriter, request *http.Request) {
			spells, err := dnd5eClient.GetSpells()
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(spells)
		})

		r.Get("/spells/{index}", func(responseWriter http.ResponseWriter, request *http.Request) {
			spell, err := dnd5eClient.GetSpell(chi.URLParam(request, "index"))
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(spell)
		})

		r.Get("/equipment", func(responseWriter http.ResponseWriter, request *http.Request) {
			equipment, err := dnd5eClient.GetEquipment()
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(equipment)
		})

		r.Get("/equipment/{index}", func(responseWriter http.ResponseWriter, request *http.Request) {
			item, err := dnd5eClient.GetEquipmentItem(chi.URLParam(request, "index"))
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadGateway)
				return
			}
			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(item)
		})
	})

	return router
}
