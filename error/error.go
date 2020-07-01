package error

import (
	"errors"
	"gorm/models"

	uuid "github.com/satori/go.uuid"
)

type ValidationError struct {
	errorMessage string
}

//NewValidationError :
func NewValidationError() ValidationError {
	return ValidationError{
		errorMessage: "",
	}
}
func sendError(err ValidationError) error {
	return errors.New(err.errorMessage)

}

func (err ValidationError) CheckUserNameError(user *models.User) error {

	if user.Name == "" {
		err.errorMessage = "User name is empty"
		return sendError(err)
	}
	return nil
}

func (err ValidationError) CheckBankNameError(bank *models.Bank) error {

	if bank.Name == "" {
		err.errorMessage = "Bank name is empty"
		return sendError(err)
	}
	return nil
}

func (err ValidationError) CheckIDError(id uuid.UUID) error {
	tmp := id.String()
	if tmp == "" {
		err.errorMessage = "ID is empty"
		return sendError(err)
	}

	return nil
}
