package service

import (
	"github.com/gin-gonic/gin"
	"go_gin_gorm_rest/db"
	"go_gin_gorm_rest/entity"
)

type Service struct {
}

type Account entity.Account

func (s Service) GetAccountAll() ([]Account, error) {
	db := db.GetDB()
	var a []Account

	if err := db.Find(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (s Service) CreateAccount(c *gin.Context) (Account, error) {
	db := db.GetDB()
	var a Account

	if err := c.BindJSON(&a); err != nil {
		return a, err
	}

	if err := db.Create(&a).Error; err != nil {
		return a, err
	}

	return a, nil
}

func (s Service) GetAccount(id string) (Account, error) {
	db := db.GetDB()
	var a Account

	if err := db.Where("id = ?", id).First(&a).Error; err != nil {
		return a, err
	}

	return a, nil
}

func (s Service) UpdateAccount(id string, c *gin.Context) (Account, error) {
	db := db.GetDB()
	var a Account

	if err := db.Where("id = ?", id).First(&a).Error; err != nil {
		return a, err
	}

	if err := c.BindJSON(&a); err != nil {
		return a, err
	}

	db.Save(&a)

	return a, nil
}

func (s Service) DeleteAccount(id string) error {
	db := db.GetDB()
	var a Account

	if err := db.Where("id = ?", id).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}
