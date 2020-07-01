package service

import (
	errors "gorm/error"
	models "gorm/models"
	"gorm/repository"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//UserAccountService :
type UserAccountService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

//NewUserAccountService :
func NewUserAccountService(db *gorm.DB, repo *repository.GormRepository) *UserAccountService {
	return &UserAccountService{
		DB:         db.AutoMigrate(&models.User{}),
		Repository: repo,
	}
}

//AddUserAccount :
func (service *UserAccountService) AddUserAccount(model *models.User) error {
	errorCheck := errors.NewValidationError()
	if err := errorCheck.CheckUserNameError(model); err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.DB, false)
	model.ID = uuid.NewV4()
	bank := &BankAccountService{
		DB:         service.DB.AutoMigrate(&models.User{}),
		Repository: service.Repository,
	}
	for i := range model.Accounts {
		model.Accounts[i].ID = uuid.NewV4()
		bank.AddBankAccount(&model.Accounts[i])

	}
	err := service.Repository.Add(uow, model)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

//GetAllUserAccounts   :Will get all accounts!
func (service *UserAccountService) GetAllUserAccounts(user *[]models.User) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	var err error
	pA := make([]string, 0)
	// var data *gorm.DB
	if err = service.Repository.GetAll(uow, &user, pA); err != nil {
		uow.Complete()
		return err
	}
	uow.Complete()
	return err
}

//DeleteUserAccount :
func (service *UserAccountService) DeleteUserAccount(id uuid.UUID) error {
	errorCheck := errors.NewValidationError()
	if err := errorCheck.CheckIDError(id); err != nil {
		return err
	}
	var user models.User
	user.ID = id
	uow := repository.NewUnitOfWork(service.DB, false)
	// pow := make([]string, 0)
	var banks []models.Bank
	pow := make([]string, 0)
	if err := service.Repository.GetAllForUserID(uow, &banks, id, pow); err != nil {
		uow.Complete()
	}
	for _, accounts := range banks {
		if err := service.Repository.Delete(uow, &accounts); err != nil {
			uow.Complete()
		}
	}
	if err := service.Repository.Delete(uow, &user); err != nil {
		uow.Complete()
	}
	uow.Commit()
	return uow.DB.Error
}

//GetUserByID : Gets user by ID
func (service *UserAccountService) GetUserByID(input *models.User, id uuid.UUID) error {
	errorCheck := errors.NewValidationError()
	if err := errorCheck.CheckUserNameError(input); err != nil {
		return err
	}
	if err := errorCheck.CheckIDError(id); err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.GetByID(uow, input, id, []string{}); err != nil {
		return err

	}
	uow.Complete()
	return nil
}

//UpdateUser :
func (service *UserAccountService) UpdateUser(input *models.User) error {
	errorCheck := errors.NewValidationError()
	if err := errorCheck.CheckUserNameError(input); err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.Update(uow, input); err != nil {
		uow.Complete()
	}
	uow.Commit()

	return uow.DB.Error
}
