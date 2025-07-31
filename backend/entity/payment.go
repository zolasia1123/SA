package entity

import (
	"time"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Payment_Date	 time.Time	`json:"start_date"`
	Amount		 float64	`json:"amount" gorm:"type:decimal(10,2)"`
	Payment_Status string  	`json:"payment_status"`

	// 1 Payment เป็นเจ้าของได้หลาย Student
	Student []Student `gorm:"foreignKey:Payment_ID"`

	// Billing_ID ทำหน้าที่เป็น FK
	Billing_ID *uint
	Billing   Billing `gorm:"foriegnKey:Billing_ID"`

}