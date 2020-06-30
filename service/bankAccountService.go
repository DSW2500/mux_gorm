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

//DeleteBankAccount :
func (service *BankAccountService) DeleteBankAccount(input interface{}) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.Delete(uow, input); err != nil {
		uow.Complete()
	}
	uow.Commit()

	return uow.DB.Error
}

//GetBankByID : Used for calling individual bank accounts through bank ID!
func (service *BankAccountService) GetBankByID(input interface{}, id interface{}) error {
	pod := make([]string, 0)
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.GetByID(uow, input, id, pod); err != nil {
		return err

	}
	uow.Complete()
	return nil
}

//GetBankByUserID : Will get all bank accounts linked by user_ID!
func (service *BankAccountService) GetBankbyUserID(input interface{}, id interface{}) error {
	pod := make([]string, 0)
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.GetByUserID(uow, input, id, pod); err != nil {
		return err

	}
	uow.Complete()
	return nil
}

//GetAllBankAccounts : Will get all accounts!
func (service *BankAccountService) GetAllBankAccounts(ba *[]models.Bank) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	var err error
	pA := make([]string, 0)
	// var data *gorm.DB
	if err = service.Repository.GetAll(uow, &ba, pA); err != nil {
		uow.Complete()
		return err
	}
	uow.Complete()
	return err
}

//UpdateBank :
func (service *BankAccountService) UpdateBank(input *models.Bank) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.Update(uow, input); err != nil {
		uow.Complete()
	}
	uow.Commit()
	uow.Complete()
	return uow.DB.Error
}