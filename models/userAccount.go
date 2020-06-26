package models

//User :
type User struct {
	Base
	Name     string
	Accounts []Bank `gorm:"ForeignKey:user_id"`
}
