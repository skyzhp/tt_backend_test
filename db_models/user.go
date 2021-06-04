package db_models

import (
	"fmt"
	//"github.com/go-pg/pg"
)

const UserType = "user"

type User struct {
	tableName struct{} `sql:"users"`
	Uid       int64
	Name      string
	Type      string
}

func CreateUser(user *User) (bool, error) {
	db := GetDB()
	_, err := db.Model(&User{Name: user.Name, Type: UserType}).Insert()
	if err != nil {
		fmt.Printf("CreateUser with err: %v user: %v  \n", err, user)
		return false, err
	}
	return true, nil
}

func GetAllUser() []*User {
	db := GetDB()
	var users []*User
	err := db.Model(&users).Order("uid").Select()
	if err != nil {
		fmt.Printf("GetAllUser with err: %v \n", err)
	}
	return users
}
