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

	// var handlers cors.Cors

	// myUser := model.User{
	// 	Name: "User2",
	// 	Accounts: []model.Bank{
	// 		{
	// 			Name:   "User2",
	// 			Amount: 9000,
	// 			Type:   "Personal"},
	// 		{
	// 			Name:   "User2",
	// 			Amount: 8900,
	// 			Type:   "Business",
	// 		},
	// 	},
	// }
	// myUser.ID = uuid.NewV4()
	// myUser.Accounts[0].ID = uuid.NewV4()
	// myUser.Accounts[1].ID = uuid.NewV4()

	//
	db, err := gorm.Open("mysql", "root:vasant12345@tcp(127.0.0.1:3306)/student?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("DB has been opened ")
	// db.CreateTable(&model.User{})
	// db.Model(&model.Bank{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE")

	// db.CreateTable(&model.Bank{})
	// db.Create(myUser)
	// db.Create(myUser2)
	// db.DropTable(&model.User{})
	//  Account details: Name : kAccount1, Amount: 25000, Type : Business
	myRepo := repository.NewRepository()
	myService := service.NewBankAccountService(db, myRepo)
	controller := controllers.NewBankAccountController(myService)
	controller.RegisterRoutes(router)
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

}

// 	fmt.Println("Choices: \n1.Create\n2.Read\n3.Update\n4.Delete\nPress 5 to quit!")
// 	flag := true
// 	if flag {
// 		fmt.Scanf("%d", &choices)
// 		switch choices {
// 		case 1:
// 			if err := myRepo.Add(uow, accHolder1); err != nil {
// 				uow.DB.Rollback()

// 			} else {
// 				fmt.Println("Added successfully!")
// 			}
// 		case 2:
// 			// Only using GetAll to avoid user input problems
// 			if err := myRepo.GetAll(uow, &accHolder2, []string{}); err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(accHolder2)
// 		case 3:

// 			fmt.Println("Updated successfully!")
// 		case 4:
// 			if err := myRepo.Delete(uow, &accHolder1); err != nil {
// 				uow.DB.Rollback()
// 				fmt.Println(err)
// 			}
// 			uow.Commit()
// 			uow.Complete()
// 			fmt.Println("Deleted successfully!")
// 		case 5:
// 			flag = false
// 			break
// 		default:
// 			fmt.Println("Wrong number!")
// 			fmt.Println("Choices: \n1.Create\n2.Read\n3.Update\n4.Delete\nPress 5 to quit!")
// 		}
// 	}

// 	// db.Update()
// 	defer db.Close()

// }

// // func checkAmount(*model.Bank *model.Bank, amount float64) error {
// // 	if *model.Bank.Amount < amount {
// // 		return errors.New("insufficient amount")
// // 	}
// // 	return nil
// // }
// // //Creating table - *model.Bank
// // db.CreateTable(&*model.Bank{})
// // account1 := &*model.Bank{
// // 	ID:     uuid.NewV4(),
// // 	Name:   "Golang",
// // 	Amount: 1500,
// // 	Type:   "Personal",
// // }
// // // Inserting account1 into db
// // db.Create(account1)
// // account2 := &*model.Bank{
// // 	ID:     uuid.NewV4(),
// // 	Name:   "SQL",
// // 	Amount: 3500,
// // 	Type:   "Business",
// // }
// // db.Create(account2)
// //db.AutoMigrate(&Customer{}, &Order{})
