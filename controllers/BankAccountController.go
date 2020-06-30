package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	model "gorm/models"
	"gorm/service"
	"io/ioutil"
	"net/http"

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

//RegisterRoutes :
func (bac *BankAccountController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/bank/account/all", bac.GetAllBankAccounts).Methods("GET")
	router.HandleFunc("/bank/user/{userID}/bankAccount", bac.GetBankAccountFromUser).Methods("GET")
	// router.HandleFunc("/bank/user/id/{id}/bankAccount/{bankID}", bac.GetBankAccountFromUser).Methods("GET")
	router.HandleFunc("/bank/user/create", bac.CreateBankData).Methods("POST")

}

//GetAllBankAccounts : Gets all the bank accounts
func (bac *BankAccountController) GetAllBankAccounts(w http.ResponseWriter, r *http.Request) {

	content := []model.Bank{}
	bac.service.GetAllBankAccounts(&content)
	// fmt.Println(content)
	RespondJSON(&w, http.StatusOK, content)

}

//GetBankAccountFromUser : gets all bank accounts for a specified user
func (bac *BankAccountController) GetBankAccountFromUser(w http.ResponseWriter, r *http.Request) {
	// file, bankID := path.Split(r.URL.String()) // extracts /bankID from the url
	// file, _ = path.Split(r.URL.String())       // removes bankAccount from url
	// file = path.Base(file)                     //extracts ID of user
	// id, _ := uuid.FromString(file)
	val := mux.Vars(r)
	id := val["userID"]
	user := model.User{}
	uas := &service.UserAccountService{
		DB:         bac.service.DB,
		Repository: bac.service.Repository,
	}
	uas.ReadByUserID(&user, id)
	// each user
	banks := []model.Bank{}
	bankUID, _ := uuid.FromString(id)
	bac.service.GetBankbyUserID(&banks, bankUID)

	RespondJSON(&w, http.StatusOK, banks)
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
