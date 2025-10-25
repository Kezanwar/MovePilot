package handlers

import (
	"formaura/pkg/output"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) (int, error) {
	defer r.Body.Close()

	return output.SuccessResponse(w, r, &output.MessageResponse{Message: "hello get users"})

}
