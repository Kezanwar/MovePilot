package handlers

import (
	user_memory_cache "movepilot/pkg/cache/user_memory"
	"movepilot/pkg/email"
	"movepilot/pkg/output"
	client_repo "movepilot/pkg/repositories/client"
	"movepilot/pkg/validate"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Request types
type CreateClientReqBody struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	AddressLine1 string   `json:"address_line1"`
	AddressLine2 string   `json:"address_line2"`
	City         string   `json:"city"`
	PostalCode   string   `json:"postal_code"`
	Country      string   `json:"country"`
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
}

func (r *CreateClientReqBody) validate() error {
	if !validate.StrNotEmpty(r.Name) {
		return fmt.Errorf("Name is required")
	}
	return nil
}

type UpdateClientReqBody struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	AddressLine1 string   `json:"address_line1"`
	AddressLine2 string   `json:"address_line2"`
	City         string   `json:"city"`
	PostalCode   string   `json:"postal_code"`
	Country      string   `json:"country"`
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
}

func (r *UpdateClientReqBody) validate() error {
	if !validate.StrNotEmpty(r.Name) {
		return fmt.Errorf("Name is required")
	}
	return nil
}

// Response types
type ClientResp struct {
	Client *client_repo.Model `json:"client"`
}

type ClientsResp struct {
	Clients []*client_repo.Model `json:"clients"`
}

type ClientHandler struct {
	ClientRepo  client_repo.Repository
	authCache   *user_memory_cache.Cache
	emailClient *email.Client
}

func NewClientHandler(
	repo client_repo.Repository,
	authCache *user_memory_cache.Cache,
	emailClient *email.Client) *ClientHandler {
	return &ClientHandler{
		ClientRepo:  repo,
		authCache:   authCache,
		emailClient: emailClient,
	}
}

func (h *ClientHandler) GetActive(w http.ResponseWriter, r *http.Request) (int, error) {
	clients, err := h.ClientRepo.FetchActive(r.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to fetch active clients")
	}

	return output.SuccessResponse(w, r, &ClientsResp{
		Clients: clients,
	})
}

func (h *ClientHandler) GetArchived(w http.ResponseWriter, r *http.Request) (int, error) {
	clients, err := h.ClientRepo.FetchArchived(r.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to fetch archived clients")
	}

	return output.SuccessResponse(w, r, &ClientsResp{
		Clients: clients,
	})
}

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	var body CreateClientReqBody

	if err := DecodeBody(r, &body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	address := client_repo.Address{
		AddressLine1: body.AddressLine1,
		AddressLine2: body.AddressLine2,
		City:         body.City,
		PostalCode:   body.PostalCode,
		Country:      body.Country,
	}

	var geo *client_repo.Geolocation
	if body.Latitude != nil && body.Longitude != nil {
		geo = &client_repo.Geolocation{
			Latitude:  *body.Latitude,
			Longitude: *body.Longitude,
		}
	}

	client, err := h.ClientRepo.Create(r.Context(), body.Name, body.Description, address, geo)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to create client")
	}

	return output.SuccessResponse(w, r, &ClientResp{
		Client: client,
	})
}

func (h *ClientHandler) View(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if !validate.StrNotEmpty(uuid) {
		return http.StatusBadRequest, fmt.Errorf("Invalid UUID")
	}

	client, err := h.ClientRepo.GetByUUID(r.Context(), uuid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to fetch client")
	}
	if client == nil {
		return http.StatusNotFound, fmt.Errorf("Client not found")
	}

	return output.SuccessResponse(w, r, &ClientResp{
		Client: client,
	})
}

func (h *ClientHandler) Update(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if !validate.StrNotEmpty(uuid) {
		return http.StatusBadRequest, fmt.Errorf("Invalid UUID")
	}

	var body UpdateClientReqBody

	if err := DecodeBody(r, &body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	address := client_repo.Address{
		AddressLine1: body.AddressLine1,
		AddressLine2: body.AddressLine2,
		City:         body.City,
		PostalCode:   body.PostalCode,
		Country:      body.Country,
	}

	var geo *client_repo.Geolocation
	if body.Latitude != nil && body.Longitude != nil {
		geo = &client_repo.Geolocation{
			Latitude:  *body.Latitude,
			Longitude: *body.Longitude,
		}
	}

	client, err := h.ClientRepo.Update(r.Context(), uuid, body.Name, body.Description, address, geo)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to update client")
	}

	return output.SuccessResponse(w, r, &ClientResp{
		Client: client,
	})
}

func (h *ClientHandler) Archive(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if !validate.StrNotEmpty(uuid) {
		return http.StatusBadRequest, fmt.Errorf("Invalid UUID")
	}

	err := h.ClientRepo.Archive(r.Context(), uuid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to archive client")
	}

	return output.SuccessResponse(w, r, nil)
}

func (h *ClientHandler) Unarchive(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if !validate.StrNotEmpty(uuid) {
		return http.StatusBadRequest, fmt.Errorf("Invalid UUID")
	}

	err := h.ClientRepo.Unarchive(r.Context(), uuid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to unarchive client")
	}

	return output.SuccessResponse(w, r, nil)
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if !validate.StrNotEmpty(uuid) {
		return http.StatusBadRequest, fmt.Errorf("Invalid UUID")
	}

	err := h.ClientRepo.SoftDelete(r.Context(), uuid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to delete client")
	}

	return output.SuccessResponse(w, r, nil)
}
