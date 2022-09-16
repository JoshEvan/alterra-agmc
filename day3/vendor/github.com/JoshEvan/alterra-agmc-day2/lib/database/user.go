package database

import (
	"github.com/JoshEvan/alterra-agmc-day3/config"
	"github.com/JoshEvan/alterra-agmc-day3/models"
)

func DeleteUserByID(id int) (err error) {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserByID(id int, updatedUser models.User) (err error) {
	if err := config.DB.Model(&updatedUser).Where("id = ?", id).Updates(updatedUser).Error; err != nil {
		return err
	}
	return
}

func AddUserByID(newUser models.User) (err error) {
	if err := config.DB.Create(&newUser).Error; err != nil {
		return err
	}
	return
}

func GetUserByID(id int) (user models.User, err error) {
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUsers() (users []models.User, err error) {
	err = config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return
}
