package controllers

import (
	"fmt"
	model "gorm/models"
	"gorm/service"
	"gorm/web"
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
	router.HandleFunc("/bank/user/bankAccount/{bankID}", bac.GetBankAccountByID).Methods("GET")
	// router.HandleFunc("/bank/user/user/{id}/bankAccount/{bankID}", bac.GetBankAccountFromUserID).Methods("GET")
	router.HandleFunc("/bank/bankAccount/create", bac.CreateBankData).Methods("POST")
	router.HandleFunc("/bank/bankAccount/delete/{bankID}", bac.DeleteBankAccount).Methods("DELETE")

}

//GetAllBankAccounts : Gets all the bank accounts
func (bac *BankAccountController) GetAllBankAccounts(w http.ResponseWriter, r *http.Request) {

	content := []model.Bank{}
	bac.service.GetAllBankAccounts(&content)
	// fmt.Println(content)
	web.RespondJSON(&w, http.StatusOK, content)

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
	uas.GetUserByID(&user, id)
	// each user
	banks := []model.Bank{}
	bankUID, _ := uuid.FromString(id)
	bac.service.GetBankbyUserID(&banks, bankUID)

	web.RespondJSON(&w, http.StatusOK, banks)
}

//GetBankAccountByID : gets all bank accounts for a specified user
func (bac *BankAccountController) GetBankAccountByID(w http.ResponseWriter, r *http.Request) {
	// file, bankID := path.Split(r.URL.String()) // extracts /bankID from the url
	// file, _ = path.Split(r.URL.String())       // removes bankAccount from url
	// file = path.Base(file)                     //extracts ID of user
	// id, _ := uuid.FromString(file)
	val := mux.Vars(r)
	id := val["bankID"]
	// each user
	banks := []model.Bank{}
	bankUID, _ := uuid.FromString(id)
	bac.service.GetBankByID(&banks, bankUID)

	web.RespondJSON(&w, http.StatusOK, banks)
}

// //GetBankAccountFromUserID : gets all bank accounts for a specified user
// func (bac *BankAccountController) GetBankAccountFromUserID(w http.ResponseWriter, r *http.Request) {
// 	// file, bankID := path.Split(r.URL.String()) // extracts /bankID from the url
// 	// file, _ = path.Split(r.URL.String())       // removes bankAccount from url
// 	// file = path.Base(file)                     //extracts ID of user
// 	// id, _ := uuid.FromString(file)
// 	val := mux.Vars(r)
// 	userID:= val["userID"]
// 	bankID:= val["bankID"]
// 	user := model.User{}
// 	uas := &service.UserAccountService{
// 		DB:         bac.service.DB,
// 		Repository: bac.service.Repository,
// 	}
// 	uas.GetUserByID(&user, userID)
// 	// each user
// 	banks := &models.Bank{}
// 	bankUID, _ := uuid.FromString(userID)
// 	bac.service.GetBankbyUserID(&banks, bankUID)
// actualBank:= &models.Bank{}
// bac.service.
// 	web.RespondJSON(&w, http.StatusOK, banks)
// }

//CreateBankData :
func (bac *BankAccountController) CreateBankData(w http.ResponseWriter, r *http.Request) {
	bank := model.Bank{}
	err := web.UnmarshalJSON(r, &bank)
	if err != nil {
		fmt.Println(err)
	}

	err = bac.service.AddBankAccount(&bank)
	if err != nil {
		fmt.Println(err)

	}
}

//DeleteBankAccount :
func (bac *BankAccountController) DeleteBankAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmp := vars["bankID"]
	id, _ := uuid.FromString(tmp)
	// user := model.User{}
	// err := UnmarshalJSON(r, &user)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	err := bac.service.DeleteBankAccount(id)
	if err != nil {
		x := []byte(err.Error())
		w.Write(x)

	}
}
