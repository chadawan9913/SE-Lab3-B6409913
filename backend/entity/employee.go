package entity

import (
	"gorm.io/gorm"
	"time"
)



type Employee struct {
	gorm.Model
	FirstName string	`gorm:"default:Employee" `
	LastName string		`gorm:"default:Employee" `
	Email string		`valid:"required~Email is required,email~Invalid email" gorm:"unique" `
	Phone string		`valid:"required~Phone is required,stringlength(10|10)~Phone must be at 10 characters" `
	Date time.Time		`valid:"required~Date is required,past~Date must be in the past" `
	Vital string		`valid:"required~Vital is required,minstringlength(5)~Vital must be at least 5 characters" `

}