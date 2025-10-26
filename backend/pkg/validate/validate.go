package validate

import (
	"regexp"
	"slices"
)

func StrNotEmpty(s ...string) bool {
	return !slices.Contains(s, "")
}

// Updated to support UUID v1-v7
var testUUID = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-7][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

func ValidateUUID(uuid string) bool {
	return testUUID.MatchString(uuid)
}

// func IsValidStatus(status string) bool {
// 	return slices.Contains(form_repo.ValidStatuses, status)
// }
