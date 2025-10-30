package handlers

import (
	user_memory_cache "movepilot/pkg/cache/user_memory"
	crm_user_repo "movepilot/pkg/repositories/crm_user"
	"movepilot/pkg/util"

	"movepilot/pkg/email"
	"movepilot/pkg/jwt"
	"movepilot/pkg/output"
	"movepilot/pkg/validate"

	"fmt"
	"net/http"
)

// Response types
type CRMManualAuthResp struct {
	User  *crm_user_repo.Model `json:"user"`
	Token string               `json:"token"`
}

type CRMAutoAuthResp struct {
	User *crm_user_repo.Model `json:"user"`
}

// Request types
type RegisterReqBody struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	TermsAndConditions bool   `json:"terms_and_conditions"`
}

func (r *RegisterReqBody) validate() error {
	if !validate.StrNotEmpty(r.FirstName, r.LastName, r.Email, r.Password) {
		return fmt.Errorf("Request body invalid")
	}
	if !r.TermsAndConditions {
		return fmt.Errorf("Terms and conditions must be accepted")
	}
	return nil
}

type SignInReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *SignInReqBody) validate() error {
	if !validate.StrNotEmpty(r.Email, r.Password) {
		return fmt.Errorf("Request body invalid")
	}
	return nil
}

type AuthHandler struct {
	CRMUserRepo crm_user_repo.Repository
	authCache   *user_memory_cache.Cache
	emailClient *email.Client
}

func NewAuthHandler(
	repo crm_user_repo.Repository,
	authCache *user_memory_cache.Cache,
	emailClient *email.Client) *AuthHandler {
	return &AuthHandler{
		CRMUserRepo: repo,
		authCache:   authCache,
		emailClient: emailClient,
	}
}

func (h *AuthHandler) CRMSignIn(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	var body SignInReqBody

	if err := DecodeBody(r, &body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	all, err := h.CRMUserRepo.FetchAll(r.Context())

	util.PrintStruct(all)

	if err != nil {
		return http.StatusBadRequest, err
	}

	usr, err := h.CRMUserRepo.GetByEmail(r.Context(), body.Email)

	fmt.Println(usr)

	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("Invalid credentials")
	}

	if !usr.IsPassword(body.Password) {
		return http.StatusBadRequest, fmt.Errorf("Invalid credentials")
	}

	tkn, err := jwt.Create(jwt.Keys.UUID, usr.UUID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to create authorization session")
	}

	return output.SuccessResponse(w, r, &CRMManualAuthResp{
		User:  usr,
		Token: tkn,
	})
}

func (h *AuthHandler) CRMInitialize(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetCRMUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	return output.SuccessResponse(w, r, &CRMAutoAuthResp{
		User: usr,
	})
}
