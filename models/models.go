package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserOperation interface {
	Create(DB *gorm.DB) (err error)
	GetOne(DB *gorm.DB, id string) (err error)
	Delete(DB *gorm.DB, id string) error

	// GetOne()
	// GetAll()
}

// Base Model for Structs
type Base struct {
	ID        uuid.UUID `gorm:"column:id; primaryKey;type:uuid; default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time `gorm:"column:created" json:"created"`
	UpdatedAt time.Time `gorm:"column:updated" json:"updated"`
}

type User struct {
	Base

	Name        string `gorm:column:name;type:string;size:256" json:"name"`
	Age         int    `gorm:column:age;type:int" json:"age"`
	Username    string `gorm:column:username;type:string;size:256" json:"username"`
	Email       string `gorm:column:email;type:string;size:256" json:"email"`
	PhoneNumber string `gorm:column:phonenumber;type:string;size:256" json:"phone_number"`
}

func (b *Base) BeforeCreate(scope *gorm.DB) error {

	b.ID = uuid.NewV4()
	return nil
}



