package repository

import (
	"auroralab/utils"
	"fmt"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Email      string `json:"email"`
	Major      string `json:"major"`
	Phone      string `json:"phone"`
	Grade      string `json:"grade"`
	Department string `json:"department"`
	IsBasics   string `json:"is_basics"`
	Introduce  string `json:"introduce"`
}

func (User) TableName() string {
	return "user"
}

var (
	user *User
)

type UserDao struct {
}

type Departments []struct {
	Department string
	Count      int64
}

func (*UserDao) AddUser(user User) error {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

func (*UserDao) SelectDe() (Departments, error) {
	var dep Departments
	result := utils.DB.Table("user").Select("department, COUNT(*) as count").Group("department").Scan(&dep)
	if result.Error != nil {
		fmt.Println(result.Error)
		return dep, result.Error
	}
	return dep, nil
}
