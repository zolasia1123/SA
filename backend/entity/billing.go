package entity

import (
	"time"
	"gorm.io/gorm"
)

type Billing struct {
	gorm.Model
	Payment_Date	time.Time	`json:"start_date"`
	Amount		 	float64		`json:"amount" gorm:"type:decimal(10,2)"`
	Payment_Status 	string  	`json:"payment_status"`

	// หลาย Billing มีได้หลาย student
	//Student []Student `gorm:"many2many:Billing_ID"`

	// 1 Billing มีได้หลาย Payment
	Payment []Payment `gorm:"foriegnKey:Billing_ID"`

	// Contract_ID ทำหน้าที่เป็น FK
	Contract_ID *uint
	Contract Contract `gorm:"foriegnKey:Contract_ID"`
}