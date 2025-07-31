package entity

import (
	"time"
	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	Start_Date		 time.Time	`json:"start_date"`
	End_Date		 time.Time	`json:"end_date"`
	Security_Deposit float64  	`json:"security_deposit" gorm:"type:decimal(10,2)"`
	Rate			 float64  	`json:"security_deposit" gorm:"type:decimal(10,2)"`

	// 1 Contract มีได้หลาย Billing
	Billing []Billing `gorm:"foreignKey:Contract_ID"`

	// Room_ID ทำหน้าที่เป็น FK
	Room_ID *uint
	Room   Room `gorm:"foriegnKey:Room_ID"`
	
	// Admin_ID ทำหน้าที่เป็น FK
	Admin_ID *uint
	Admin Admin `gorm:"foriegnKey:Admin_ID"`
}