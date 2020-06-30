package service

import (
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
	return err
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
	var user models.User
	user.ID = id
	uow := repository.NewUnitOfWork(service.DB, false)
	// pow := make([]string, 0)
	var banks []models.Bank
	if err := service.ReadByUserID(&banks, id); err != nil {
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

//ReadByUserID : Will get all accounts related to specified user
func (service *UserAccountService) ReadByUserID(input interface{}, id interface{}) error {

	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.GetAllForUsersID(uow, input, id, []string{}); err != nil {
		return err

	}
	uow.Complete()
	return nil
}

//UpdateUser :
func (service *UserAccountService) UpdateUser(input *models.User) error {
	uow := repository.NewUnitOfWork(service.DB, false)
	if err := service.Repository.Update(uow, input); err != nil {
		uow.Complete()
	}
	uow.Commit()
	uow.Complete()
	return uow.DB.Error
}
