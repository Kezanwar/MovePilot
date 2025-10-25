package handlers_test

import (
	"move-pilot/cmd/api/handlers"
	"move-pilotot/pkg/output"
	user_repo "move-pilotot/pkg/repositories/user"
	"move-pilotot/pkg/util"

	"context"
	"net/http"
	"testing"
)

type mockUserRepo struct {
	user_repo.Repository // ✅ embed the interface (optional, for clarity/logging)
	CreateFn             func(ctx context.Context, firstName, lastName, email, password, otp string, termsAndConditions bool) (*user_repo.Model, error)
	DoesEmailExistFn     func(ctx context.Context, email string) (bool, error)
}

func (m *mockUserRepo) Create(ctx context.Context, firstName, lastName, email, password, otp string, termsAndConditions bool) (*user_repo.Model, error) {
	return m.CreateFn(ctx, firstName, lastName, email, password, otp, termsAndConditions)
}

func (m *mockUserRepo) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	return m.DoesEmailExistFn(ctx, email)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := &mockUserRepo{
		DoesEmailExistFn: func(ctx context.Context, email string) (bool, error) {
			return false, nil
		},
		CreateFn: func(ctx context.Context, firstName, lastName, email, password, otp string, termsAndConditions bool) (*user_repo.Model, error) {

			return &user_repo.Model{
				UUID:               "test-uuid",
				FirstName:          firstName,
				LastName:           lastName,
				Email:              email,
				OTP:                otp,
				TermsAndConditions: termsAndConditions,
			}, nil
		},
	}
	// inside TestRegister_Success
	handler := handlers.NewAuthHandler(mockRepo)
	wrapped := output.MakeJsonHandler(handler.Register)

	body := map[string]interface{}{
		"first_name":           "Test",
		"last_name":            "User",
		"email":                "test@example.com",
		"password":             "secure123",
		"terms_and_conditions": true,
	}

	res, status := util.TestJsonRequestAndDecode[handlers.ManualAuthResp](t, wrapped, http.MethodPost, "/api/auth/register", body)

	if status != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", status)
	}

	if res.User.Email != "test@example.com" {
		t.Errorf("expected email %q, got %q", "test@example.com", res.User.Email)
	}

	if res.Token == "" {
		t.Error("expected token to be set")
	}
}
