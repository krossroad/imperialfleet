package imperialfleet

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/krossroad/imperialfleet/internal/logger"
	"github.com/krossroad/imperialfleet/internal/models"
	"github.com/krossroad/imperialfleet/internal/persist"
)

func New(log *logger.Entry, p persist.Persist) *Service {
	return &Service{
		persist:  p,
		log:      log,
		validate: validator.New(),
	}
}

type Service struct {
	log      *logger.Entry
	persist  persist.Persist
	validate *validator.Validate
}

func (s *Service) List(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	enc := json.NewEncoder(res)

	lr := models.ListCraftRequest{
		Name:   mux.Vars(req)["name"],
		Class:  mux.Vars(req)["class"],
		Status: mux.Vars(req)["status"],
	}

	items, err := s.persist.List(ctx, lr)
	if err != nil {
		s.log.Error("failed to list items", "error", err)
		res.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]any{
			"error": "unable to list items",
		})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(items); err != nil {
		s.log.Error("failed to encode response", "error", err)
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) Create(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	enc := json.NewEncoder(res)
	var craft models.SpaceCraft
	if err := json.NewDecoder(req.Body).Decode(&craft); err != nil {
		s.log.Error("failed to decode request", "error", err)
		res.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]any{
			"error": "invalid request",
		})
		return
	}

	if failed := s.validationHandler(craft, res, enc); failed {
		return
	}

	craft.CreatedAt = time.Now()
	craft.UpdatedAt = time.Now()

	if err := s.persist.Create(ctx, &craft); err != nil {
		s.log.Error("failed to create craft", "error", err)
		res.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]any{
			"error": "unable to create craft",
		})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc.Encode(map[string]any{
		"success": true,
	})
}

func (s *Service) Delete(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	reqID := mux.Vars(req)["id"]
	enc := json.NewEncoder(res)

	id, err := strconv.Atoi(reqID)
	if err != nil {
		s.log.Error("failed to parse ID", "error", err)
		res.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]any{
			"error": "invalid ID",
		})
	}

	if err := s.persist.Delete(ctx, id); err != nil {
		s.log.Error("failed to delete craft", "error", err)
		res.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]any{
			"error": "unable to delete craft",
		})
		return
	}

	enc.Encode(map[string]any{
		"success": true,
	})
}

func (s *Service) Update(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	reqID := mux.Vars(req)["id"]
	enc := json.NewEncoder(res)

	id, err := strconv.Atoi(reqID)
	if err != nil {
		s.log.Error("failed to parse ID", "error", err)
		res.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]any{
			"error": "invalid ID",
		})
		return
	}

	var craft models.SpaceCraft
	if err := json.NewDecoder(req.Body).Decode(&craft); err != nil {
		s.log.Error("failed to decode request", "error", err)
		res.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]any{
			"error": "invalid request",
		})
		return
	}
	craft.ID = id
	craft.UpdatedAt = time.Now()

	if failed := s.validationHandler(craft, res, enc); failed {
		return
	}

	if err := s.persist.Update(ctx, &craft); err != nil {
		s.log.Error("failed to update craft", "error", err)
		res.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]any{
			"error": "unable to update craft",
		})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	enc.Encode(map[string]any{
		"success": true,
	})
}

func (s *Service) validationHandler(craft models.SpaceCraft, res http.ResponseWriter, enc *json.Encoder) bool {
	if err := s.validate.Struct(&craft); err != nil {
		if vErr, ok := err.(*validator.ValidationErrors); ok {
			s.log.Error("failed to validate request", "error", vErr.Error())
			res.WriteHeader(http.StatusBadRequest)
			enc.Encode(map[string]any{
				"error": "failed validation: " + vErr.Error(),
			})
			return true
		}

		s.log.Error("failed to validate request", "error", err)
		res.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]any{
			"error": "invalid request body",
		})
		return true
	}
	return false
}

// Show
func (s *Service) Show(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	reqID := mux.Vars(req)["id"]
	enc := json.NewEncoder(res)

	id, err := strconv.Atoi(reqID)
	if err != nil {
		s.log.Error("failed to parse ID", "error", err)
		res.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]any{
			"error": "invalid ID",
		})
		return
	}

	craft, err := s.persist.Get(ctx, id)
	if err != nil {
		s.log.Error("failed to show craft", "error", err)
		res.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]any{
			"error": "unable to show craft",
		})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	enc.Encode(craft)
}
