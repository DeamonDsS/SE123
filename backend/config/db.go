package config


import (

   "fmt"

   "time"

   "github.com/SE67/entity"

   "gorm.io/driver/sqlite"

   "gorm.io/gorm"

)


var db *gorm.DB


func DB() *gorm.DB {

   return db

}


func ConnectionDB() {

   database, err := gorm.Open(sqlite.Open("se67.db?cache=shared"), &gorm.Config{})

   if err != nil {

       panic("failed to connect database")

   }

   fmt.Println("connected database")

   db = database

}


func SetupDatabase() {


   db.AutoMigrate(

       &entity.Users{},
       &entity.Genders{},
       &entity.Events{},
       &entity.Locations{},
       &entity.Path{},
       &entity.Timers{},
       &entity.TpyeEvents{},
       &entity.Order{},
       &entity.Payment{},
       &entity.Ticket{},
       &entity.Tpackage{},
       &entity.Code{},

   )


   GenderMale := entity.Genders{Gender: "Male"}

   GenderFemale := entity.Genders{Gender: "Female"}


   db.FirstOrCreate(&GenderMale, &entity.Genders{Gender: "Male"})

   db.FirstOrCreate(&GenderFemale, &entity.Genders{Gender: "Female"})


   hashedPassword, _ := HashPassword("123456")
   BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")

   User := &entity.Users{

       FirstName: "Admin",

       LastName:  "Admin",

       Email:     "                                     ",

       Age:       10,

       Password:  hashedPassword,

       BirthDay:  BirthDay,

       GenderID:  1,
       IsAdmin:   true,

   }

   db.FirstOrCreate(User, &entity.Users{

       Email: "Admin@gmail.com",

   })

  
}