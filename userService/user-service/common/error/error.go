package errorhandler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidationResponse(err error) (responses []ValidationResponse) {
	var fieldErrors validator.ValidationErrors

	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				responses = append(responses, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is required", err.Field()),
				})
			case "email":
				responses = append(responses, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is not a valid email address", err.Field()),
				})
			default:
				errValidator, ok := ErrValidator[err.Tag()]
				if ok {
					count := strings.Count(errValidator, "%s")
					if count == 1 {
						responses = append(responses, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field()),
						})
					} else {
						responses = append(responses, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field(), err.Param()),
						})
					}
				} else {
					responses = append(responses, ValidationResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag()),
					})
				}
			}
		}
	}

	return responses
}

func WrapError(err error) error {
	logrus.Errorf("error: %v", err)
	return err
}
