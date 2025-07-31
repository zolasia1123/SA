package config


import (

   "fmt"

   "github.com/SA/entity"

   "gorm.io/driver/sqlite"

   "gorm.io/gorm"

   "time"

)


var db *gorm.DB


func DB() *gorm.DB {

   return db

}


func ConnectionDB() {

   database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})

   if err != nil {

       panic("failed to connect database")

   }

   fmt.Println("connected database")

   db = database

}

func SetupDatabase() {
   db.AutoMigrate(&entity.Admin{}, &entity.Billing{}, &entity.Payment{}, &entity.Contract{}, &entity.Room{}, &entity.Student{})


   hashedPassword, _ := HashPassword("123456")
   BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")
   User := &entity.Student{
      First_Name: "Software",
      Last_Name:  "Analysis",
       Email:     "sa@gmail.com",
       Password:  hashedPassword,
       Birthday:  BirthDay,
   }
   db.FirstOrCreate(User, &entity.Student{
       Email: "sa@gmail.com",
   })

   startDate, _ := time.Parse("2006-01-02", "2025-08-01")
	endDate, _ := time.Parse("2006-01-02", "2026-07-31")
   // สร้าง Contract ใหม่
   Contract := &entity.Contract{
		Start_Date:       startDate,
		End_Date:         endDate,
		Security_Deposit: 5000.00,
		Rate:             7500.00,
	}
   db.FirstOrCreate(Contract, &entity.Contract{})
}

