package service

import (
	models "gorm/models"
	"gorm/repository"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//BankAccountService :
type BankAccountService struct {
	DB         *gorm.DB
	Repository *repository.GormRepository
}

//NewBankAccountService :
func NewBankAccountService(db *gorm.DB, repository *repository.GormRepository) *BankAccountService {
	db = db.AutoMigrate(&models.Bank{})
	return &BankAccountService{
		DB:         db,
		Repository: repository,
	}

}

//AddUserAccount :
func (service *BankAccountService) AddUserAccount(model *models.User) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	model.ID = uuid.NewV4()
	for i := range model.Accounts {
		model.Accounts[i].ID = uuid.NewV4()
		service.AddBankAccount(&model.Accounts[i])
	}
	err := service.Repository.Add(uow, model)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//AddBankAccount :
func (service *BankAccountService) AddBankAccount(model *models.Bank) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	model.ID = uuid.NewV4()
	err := service.Repository.Add(uow, model)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

//Delete :
func (service *BankAccountService) DeleteBankAccount(input *models.Bank) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.Delete(uow, &input); err != nil {
		uow.Complete()
	}
	uow.Commit()
	uow.Complete()
	return uow.DB.Error
}

//ReadByID : Used for calling individual bank accounts!
func (service *BankAccountService) ReadByID(input interface{}, id uuid.UUID) error {
	pod := make([]string, 0)
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.GetByID(uow, input, id, pod); err != nil {
		return err
		uow.Complete()

	}
	uow.Complete()
	return nil
}

//ReadByUserID : Will get all accounts related to specified user
func (service *BankAccountService) ReadByUserID(input interface{}, id interface{}) error {

	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.GetAllForUsersID(uow, input, id, []string{}); err != nil {
		return err
		uow.Complete()

	}
	uow.Complete()
	return nil
}

//GetAllAccounts : Will get all accounts!
func (service *BankAccountService) GetAllAccounts(ba interface{}) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	var err error
	pA := make([]string, 0)
	// var data *gorm.DB
	if err = service.Repository.GetAll(uow, ba, pA); err != nil {
		uow.Complete()
		return err
	}
	uow.Complete()
	return err
}

//Update :
func (service *BankAccountService) Update(input *models.Bank) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.Update(uow, input); err != nil {
		uow.Complete()
	}
	uow.Commit()
	uow.Complete()
	return uow.DB.Error
}