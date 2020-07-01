package controllers

import (
	"gorm/models"

	"gorm/service"
	"gorm/web"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

//UserAccountController :
type UserAccountController struct {
	service *service.UserAccountService
}

//NewUserAccountController :
func NewUserAccountController(uas *service.UserAccountService) *UserAccountController {
	return &UserAccountController{
		service: uas,
	}
}

//RegisterRoutes :
func (uas *UserAccountController) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/bank/user/all", uas.GetAllUsers).Methods("GET")
	router.HandleFunc("/bank/user/{id}", uas.GetUserByID).Methods("GET")
	router.HandleFunc("/bank/user/create", uas.CreateUserData).Methods("POST")
	router.HandleFunc("/bank/user/update", uas.UpdateUserData).Methods("PUT")
	router.HandleFunc("/bank/user/delete/{userID}", uas.DeleteUser).Methods("DELETE")

}

//CreateUserData :
func (uas *UserAccountController) CreateUserData(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		x := []byte(err.Error())
		w.Write(x)
	}

	err = uas.service.AddUserAccount(&user)
	if err != nil {
		x := []byte(err.Error())
		w.Write(x)

	}
}

//UpdateUserData :
func (uas *UserAccountController) UpdateUserData(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {

		x := []byte(err.Error())
		w.Write(x)
	}

	err = uas.service.UpdateUser(&user)
	if err != nil {
		x := []byte(err.Error())
		w.Write(x)
	}

}

//DeleteUser :
func (uas *UserAccountController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmp := vars["userID"]
	id, _ := uuid.FromString(tmp)
	// user := model.User{}
	// err := UnmarshalJSON(r, &user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err := uas.service.DeleteUserAccount(id)
	if err != nil {
		x := []byte(err.Error())
		w.Write(x)

	}
}

//GetAllUsers : Gets all the users available
func (uas *UserAccountController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}

	uas.service.GetAllUserAccounts(&users)
	for i := range users { // each user
		banks := []models.Bank{}

		bac := &service.BankAccountService{
			DB:         uas.service.DB,
			Repository: uas.service.Repository,
		}

		bac.GetBankbyUserID(&banks, users[i].ID)
		for _, val := range banks {
			users[i].Accounts = append(users[i].Accounts, val)
		}
	}
	web.RespondJSON(&w, http.StatusOK, users)
}

//GetUserByID : gets a user by specified id
func (uas *UserAccountController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// path := path.Base(r.URL.String())
	// id, _ := uuid.FromString(path)
	val := mux.Vars(r)
	id := val["id"]
	user := models.User{}
	uas.service.GetUserByID(&user, id)
	// each user
	var banks []models.Bank

	bac := &service.BankAccountService{
		DB:         uas.service.DB,
		Repository: uas.service.Repository,
	}

	bac.GetBankbyUserID(&banks, user.ID)
	for _, val := range banks {
		user.Accounts = append(user.Accounts, val)
	}

	web.RespondJSON(&w, http.StatusOK, user)
}
