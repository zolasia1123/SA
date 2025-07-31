package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Username 	 string 	`json:"username"`
	Password 	 string		`json:"password"`
	Email	 	 string		`json:"email"`
	First_Name 	 string		`json:"first_name"`
	Last_Name    string  	`json:"last_name"`
	Birthday	 time.Time	`json:"birthday"`
	Phone		 string		`json:"phone"`
	Parent_Phone string		`json:"parent_phone"`
	Parent_Name  string		`json:"parent_name"`

	// หลาย student มีได้หลาย Billing
	Billing []Billing `gorm:"many2many:Student_ID"`

	// MemberID ทำหน้าที่เป็น FK
	Room_ID *uint
	Room   Room `gorm:"foriegnKey:Room_ID"`

	// Contract_ID ทำหน้าที่เป็น FK
	Contract_ID *uint
	Contract Contract `gorm:"foriegnKey:Contract_ID"`

	// Payment_ID ทำหน้าที่เป็น FK
	Payment_ID *uint
	Payment Payment `gorm:"foriegnKey:Payment_ID"`
}
