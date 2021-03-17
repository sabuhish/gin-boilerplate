package repository

import (
	"errors"
	"fmt"
	"gin-boilerplate/database"
	"gin-boilerplate/models"

	"github.com/gofrs/uuid"
)

type UserRepository interface {
	Create(u *models.User) (*models.User, error)
	GetOne(u *models.User, id string) error
	Delete(u *models.User, id string) error
	Update(u *models.User, id string) error
}

type DBUserRepositoty struct{}

func (o *DBUserRepositoty) Create(u *models.User) (*models.User, error) {

	result := database.DB.Create(&u)

	if result.Error != nil {
		fmt.Println(result.Error)
		return u, errors.New("Cannot crete user")
	}
	return u, nil
}

func (o *DBUserRepositoty) GetOne(u *models.User, id string) error {

	uid, _ := uuid.FromString(id)

	if err := database.DB.Where("id = ?", uid).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (o *DBUserRepositoty) Delete(u *models.User, id string) error {

	err := o.GetOne(u, id)

	if err != nil {
		return err
	}
	database.DB.Delete(&u)

	return nil

}

func (o *DBUserRepositoty) Update(u *models.User, id string) error {

	err := o.GetOne(u, id)
	if err != nil {
		return err
	}

	database.DB.Save(&u)

	return nil
}
