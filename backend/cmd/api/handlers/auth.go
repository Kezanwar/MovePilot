package handlers

import (
	user_memory_cache "movepilot/pkg/cache/user_memory"
	user_repo "movepilot/pkg/repositories/user"

	"movepilot/pkg/email"
	"movepilot/pkg/jwt"
	"movepilot/pkg/otp"
	"movepilot/pkg/output"
	"movepilot/pkg/validate"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Response types
type ManualAuthResp struct {
	User  *user_repo.Model `json:"user"`
	Token string           `json:"token"`
}

type AutoAuthResp struct {
	User *user_repo.Model `json:"user"`
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
	UserRepo    user_repo.Repository
	authCache   *user_memory_cache.Cache
	emailClient *email.Client
}

func NewAuthHandler(
	repo user_repo.Repository,
	authCache *user_memory_cache.Cache,
	emailClient *email.Client) *AuthHandler {
	return &AuthHandler{
		UserRepo:    repo,
		authCache:   authCache,
		emailClient: emailClient,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	var body RegisterReqBody
	if err := DecodeBody(r, &body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	exists, err := h.UserRepo.DoesEmailExist(r.Context(), body.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if exists {
		return http.StatusBadRequest, fmt.Errorf("This email already exists")
	}

	// Generate OTP for email verification
	userOTP, err := otp.Generate()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to generate OTP: %w", err)
	}

	usr, err := h.UserRepo.Create(r.Context(), body.FirstName, body.LastName, body.Email, body.Password, userOTP, body.TermsAndConditions)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// TODO: Send OTP email to user
	// err = h.emailClient.SendOTP(email.OTPEmailData{
	// 	ToEmail: usr.Email,
	// 	ToName:  fmt.Sprintf("%s %s", usr.FirstName, usr.LastName),
	// 	OTPCode: userOTP,
	// })
	// if err != nil {
	// 	return http.StatusInternalServerError, fmt.Errorf("failed to send OTP email: %w", err)
	// }

	// TODO: rmv this
	fmt.Println(userOTP)

	tkn, err := jwt.Create(jwt.Keys.UUID, usr.UUID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return output.SuccessResponse(w, r, &ManualAuthResp{
		User:  usr,
		Token: tkn,
	})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	var body SignInReqBody

	if err := DecodeBody(r, &body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	usr, err := h.UserRepo.GetByEmail(r.Context(), body.Email)

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

	return output.SuccessResponse(w, r, &ManualAuthResp{
		User:  usr,
		Token: tkn,
	})
}

func (h *AuthHandler) Initialize(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	fmt.Println(usr.OTP)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	return output.SuccessResponse(w, r, &AutoAuthResp{
		User: usr,
	})
}

func (h *AuthHandler) ConfirmOTP(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	vars := mux.Vars(r)
	otpParam := vars["otp"]
	if otpParam == "" {
		return http.StatusBadRequest, fmt.Errorf("OTP parameter is required")
	}

	if !usr.ValidateOTP(otpParam) {
		return http.StatusBadRequest, fmt.Errorf("Invalid OTP")
	}

	usr.EmailConfirmed = true
	err = h.UserRepo.UpdateEmailConfirmed(r.Context(), usr.UUID, true)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Failed to update email confirmation status")
	}

	h.authCache.Delete(usr.UUID)

	return output.SuccessResponse(w, r, &AutoAuthResp{
		User: usr,
	})
}

func (h *AuthHandler) ResendOTP(w http.ResponseWriter, r *http.Request) (int, error) {

	usr, err := GetUserFromCtx(r)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	if usr.EmailConfirmed {
		return http.StatusBadRequest, fmt.Errorf("Email already confirmed")
	}

	newOTP, err := otp.Generate()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to generate OTP: %w", err)
	}

	err = h.UserRepo.UpdateOTP(r.Context(), usr.UUID, newOTP)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update OTP: %w", err)
	}

	// TODO: rmv this
	fmt.Println(newOTP)

	// TODO: bring back email
	// // Send new OTP email
	// err = h.emailClient.SendOTP(email.OTPEmailData{
	// 	ToEmail: usr.Email,
	// 	ToName:  fmt.Sprintf("%s %s", usr.FirstName, usr.LastName),
	// 	OTPCode: newOTP,
	// })
	// if err != nil {
	// 	return http.StatusInternalServerError, fmt.Errorf("failed to send OTP email: %w", err)
	// }

	// Clear cache to force fresh user data on next request
	h.authCache.Delete(usr.UUID)

	return output.SuccessResponse(w, r, map[string]string{
		"message": "OTP resent successfully",
	})
}
