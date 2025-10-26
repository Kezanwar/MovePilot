package handlers

import (
	"encoding/json"
	"fmt"
	"movepilot/pkg/constants"
	user_repo "movepilot/pkg/repositories/user"
	"movepilot/pkg/validate"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserFromCtx(r *http.Request) (*user_repo.Model, error) {
	usr, ok := r.Context().Value(constants.USER_CTX).(*user_repo.Model)

	if !ok {
		return nil, fmt.Errorf("handlers.GetUserFromContext: cant find user in r.Context")
	}

	return usr, nil

}

func GetUUIDFromParams(r *http.Request) (*string, error) {
	vars := mux.Vars(r)
	formUuid := vars["uuid"]
	if formUuid == "" {
		return nil, fmt.Errorf("UUID is required")
	}

	if !validate.ValidateUUID(formUuid) {
		return nil, fmt.Errorf("UUID is required")
	}

	return &formUuid, nil

}

func DecodeBody(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
