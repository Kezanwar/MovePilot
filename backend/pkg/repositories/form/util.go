package form_repo

import (
	"fmt"
	"regexp"
	"strconv"
)

var untitledPattern = regexp.MustCompile(`^Untitled(?: (\d+))?$`)

func GenerateFormUntitledName(forms []*FormModel) string {
	if len(forms) == 0 {
		return "Untitled"
	}

	// Find highest number in "Untitled X" pattern
	maxNumber := 0

	for _, form := range forms {
		matches := untitledPattern.FindStringSubmatch(form.Name)
		if matches != nil {
			if matches[1] == "" {
				// Just "Untitled" with no number = Untitled 1
				if maxNumber < 1 {
					maxNumber = 1
				}
			} else {
				// "Untitled X"
				num, _ := strconv.Atoi(matches[1])
				if num > maxNumber {
					maxNumber = num
				}
			}
		}
	}

	// Generate next number
	if maxNumber == 0 {
		return "Untitled"
	}
	return fmt.Sprintf("Untitled %d", maxNumber+1)
}
