package handlers

import (
	"context"
	"fmt"
	"move-pilot/pkg/email"
	"move-pilotot/pkg/output"
	form_repo "move-pilotot/pkg/repositories/form"
	"move-pilotot/pkg/validate"
	"net/http"
)

type SubmissionHandler struct {
	FormRepo    form_repo.Repository
	emailClient *email.Client
}

func NewSubmissionHandler(
	repo form_repo.Repository,
	emailClient *email.Client) *SubmissionHandler {
	return &SubmissionHandler{
		FormRepo:    repo,
		emailClient: emailClient,
	}
}

func (h *SubmissionHandler) GetForm(w http.ResponseWriter, r *http.Request) (int, error) {

	formUuid, err := GetUUIDFromParams(r)

	if err != nil {
		return http.StatusBadRequest, err
	}

	form, err := h.FormRepo.GetByUUID(r.Context(), *formUuid)

	if err != nil {
		return http.StatusNotFound, fmt.Errorf("Resource not found")
	}

	if form.Status != form_repo.StatusActive {
		return http.StatusForbidden, fmt.Errorf("This form is currently unavailable")
	}

	err = h.FormRepo.IncrementViews(context.Background(), *formUuid)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to view form, please try again later")
	}

	return output.SuccessResponse(w, r, &GetFormResponse{
		Form: form,
	})
}

type SubmitFormReqBody struct {
	AffiliateUUID string `json:"affiliate_uuid"`
}

func (r *SubmitFormReqBody) validate() error {
	if validate.StrNotEmpty(r.AffiliateUUID) {
		if !validate.ValidateUUID(r.AffiliateUUID) {
			return fmt.Errorf("Incorrect affiliate uuid format")
		}
	}

	return nil
}

func (h *SubmissionHandler) SubmitForm(w http.ResponseWriter, r *http.Request) (int, error) {

	formUuid, err := GetUUIDFromParams(r)

	if err != nil {
		return http.StatusBadRequest, err
	}

	form, err := h.FormRepo.GetByUUID(r.Context(), *formUuid)

	if err != nil {
		return http.StatusNotFound, fmt.Errorf("Resource not found")
	}

	var body SubmitFormReqBody

	if err := DecodeBody(r, &body); err != nil {
		return http.StatusBadRequest, err
	}

	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	//TODO: Fill out rest of this handler, create a submission etc

	return output.SuccessResponse(w, r, &GetFormResponse{
		Form: form,
	})
}
