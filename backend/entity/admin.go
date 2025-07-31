package entity

import (
	"time"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username 	string 		`json:"username"`
	Password 	string		`json:"password"`
	Email	 	string		`json:"email"`
	First_Name 	string		`json:"first_name"`
	Last_Name   string  	`json:"last_name"`
	Birthday	time.Time	`json:"birthday"`
	Phone		string		`json:"phone"`

	// หลาย room อยู่ได้หลาย admin
	//Room []Room `gorm:"many2many:Admin_ID"`

	// หลาย admin อยู่ได้หลาย Contract
	Contract []Contract `gorm:"foreignKey:Admin_ID"`
}