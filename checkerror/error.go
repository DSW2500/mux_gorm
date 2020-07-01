package checkerror

import (
	"errors"
	"gorm/models"

	uuid "github.com/satori/go.uuid"
)

type errortype struct {
	errorMessage string
}

func NewErrorType() errortype {
	return errortype{
		errorMessage: "",
	}
}
func sendError(err errortype) error {
	return errors.New(err.errorMessage)

}

func (err errortype) CheckUserNameError(user *models.User) error {

	if user.Name == "" {
		err.errorMessage = "User name is empty"
		return sendError(err)
	}
	return nil
}

func (err errortype) CheckBankNameError(bank *models.Bank) error {

	if bank.Name == "" {
		err.errorMessage = "Bank name is empty"
		return sendError(err)
	}
	return nil
}

func (err errortype) CheckIDError(id uuid.UUID) error {
	tmp := id.String()
	if tmp == "" {
		err.errorMessage = "ID is empty"
		return sendError(err)
	}

	return nil
}
