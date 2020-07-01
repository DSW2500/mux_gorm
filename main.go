package main

import (
	"fmt"
	"gorm/controllers"
	"gorm/repository"
	"gorm/service"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "database/sql"
)

func main() {

	// Creating a router
	router := mux.NewRouter()

	db, err := gorm.Open("mysql", "root:vasant12345@tcp(127.0.0.1:3306)/student?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("DB has been opened ")

	myRepo := repository.NewRepository()
	myUserService := service.NewUserAccountService(db, myRepo)
	myBankService := service.NewBankAccountService(db, myRepo)
	UserAccountController := controllers.NewUserAccountController(myUserService)
	BankAccountController := controllers.NewBankAccountController(myBankService)
	UserAccountController.RegisterRoutes(router)
	BankAccountController.RegisterRoutes(router)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "token"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	srv := &http.Server{
		Handler: handlers.CORS(headers, methods, origin)(router),
		Addr:    ":8800",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	fmt.Println("The service is working!")
	// db.AutoMigrate(&models.User{}, &models.Bank{})
}
