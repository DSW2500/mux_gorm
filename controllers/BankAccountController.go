package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	model "gorm/models"
	"gorm/service"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

//BankAccountController :
type BankAccountController struct {
	service *service.BankAccountService
}

//NewBankAccountController :
func NewBankAccountController(service *service.BankAccountService) *BankAccountController {
	return &BankAccountController{
		service: service,
	}
}

//GetAllBankAccounts : Gets all the bank accounts
func (bac *BankAccountController) GetAllBankAccounts(w http.ResponseWriter, r *http.Request) {

	content := []model.Bank{}
	bac.service.GetAllAccounts(&content)
	// fmt.Println(content)
	RespondJSON(&w, http.StatusOK, content)

}

// GetAllBankAccounts is working properly!

/////////////////////

//GetAllUsers : Gets all the users available
func (bac *BankAccountController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []model.User{}

	bac.service.GetAllAccounts((&users))
	for i := range users { // each user
		banks := []model.Bank{}
		bac.service.ReadByUserID(&banks, users[i].ID)
		for _, val := range banks {
			users[i].Accounts = append(users[i].Accounts, val)
		}
	}
	RespondJSON(&w, http.StatusOK, users)
}

//GetAllUsers is working properly!
////////////////////////////

//GetUserByID : gets a user by specified id
func (bac *BankAccountController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	path := path.Base(r.URL.String())
	id, _ := uuid.FromString(path)
	user := model.User{}
	bac.service.ReadByID(&user, id)
	// each user
	banks := []model.Bank{}
	bac.service.ReadByUserID(&banks, user.ID)
	for _, val := range banks {
		user.Accounts = append(user.Accounts, val)
	}

	RespondJSON(&w, http.StatusOK, user)
}

//GetUserIDBankAccounts : gets all bank accounts for a specified user
func (bac *BankAccountController) GetUserIDBankAccounts(w http.ResponseWriter, r *http.Request) {
	file, _ := path.Split(r.URL.String()) // removes /bankAccount from the url
	file = path.Base(file)                //extracts ID of user
	id, _ := uuid.FromString(file)
	user := model.User{}
	bac.service.ReadByID(&user, id)
	// each user
	banks := []model.Bank{}
	bac.service.ReadByUserID(&banks, user.ID)
	for _, val := range banks {
		user.Accounts = append(user.Accounts, val)
	}
	RespondJSON(&w, http.StatusOK, banks)
}

//GetBankAccountFromUser : gets all bank accounts for a specified user
func (bac *BankAccountController) GetBankAccountFromUser(w http.ResponseWriter, r *http.Request) {
	file, bankID := path.Split(r.URL.String()) // extracts /bankID from the url
	file, _ = path.Split(r.URL.String())       // removes bankAccount from url
	file = path.Base(file)                     //extracts ID of user
	id, _ := uuid.FromString(file)
	user := model.User{}
	bac.service.ReadByID(&user, id)
	// each user
	banks := []model.Bank{}
	bac.service.ReadByUserID(&banks, user.ID) //Gets all the banks from the specified user
	bankUID, _ := uuid.FromString(bankID)
	bac.service.ReadByID(&banks, bankUID)

	RespondJSON(&w, http.StatusOK, banks)
}

//CreateUserData :
func (bac *BankAccountController) CreateUserData(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := UnmarshalJSON(r, &user)
	if err != nil {
		fmt.Println(err)
	}

	err = bac.service.AddUserAccount(&user)
	if err != nil {
		fmt.Println(err)

	}
}

//CreateBankData :
func (bac *BankAccountController) CreateBankData(w http.ResponseWriter, r *http.Request) {
	bank := model.Bank{}
	err := UnmarshalJSON(r, &bank)
	if err != nil {
		fmt.Println(err)
	}

	err = bac.service.AddBankAccount(&bank)
	if err != nil {
		fmt.Println(err)

	}
}

//UnmarshalJSON :
func UnmarshalJSON(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return errors.New("There is problem while reading data")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("Can't handle data")
	}

	if len(body) == 0 {
		return errors.New("Empty Data")
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.New("Unable to Parse Data")
	}
	return nil
}

//RegisterRoutes :
func (bac *BankAccountController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/bank/account/all", bac.GetAllBankAccounts).Methods("GET")
	router.HandleFunc("/bank/user/all", bac.GetAllUsers).Methods("GET")
	router.HandleFunc("/bank/user/id/{id}", bac.GetUserByID).Methods("GET")
	router.HandleFunc("/bank/user/id/{id}/bankAccount", bac.GetUserIDBankAccounts).Methods("GET")
	router.HandleFunc("/bank/user/id/{id}/bankAccount/{bankID}", bac.GetBankAccountFromUser).Methods("GET")
	router.HandleFunc("/bank/user/create", bac.CreateUserData).Methods("POST")
	router.HandleFunc("/bank/user/bank/create", bac.CreateBankData).Methods("POST")
}

//write to header func
func writeToHeader(w *http.ResponseWriter, statusCode int, payload interface{}) {
	(*w).WriteHeader(statusCode)
	(*w).Write(payload.([]byte))
}

//RespondJSON :
func RespondJSON(w *http.ResponseWriter, statusCode int, content interface{}) {
	response, err := json.Marshal(content)
	if err != nil {
		writeToHeader(w, http.StatusInternalServerError, err.Error())
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	writeToHeader(w, statusCode, response)
}
