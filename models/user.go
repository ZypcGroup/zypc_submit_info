package models

import (
// "fmt"
// _ "github.com/go-sql-driver/mysql"
// "github.com/go-xorm/xorm"
// "time"
// "zypc_submit/modules"
)

func AddUser(user *User) (err error) {
	connectDB()
	_, err = engine.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func CheckUser(userid int64) (has bool, err error) {
	connectDB()
	user := &User{UserId: userid}
	has, err = engine.Get(user)
	if err != nil {
		return false, err
	}
	return has, nil
}

func ModifyUser() {

}

func JudgeUser(user *User) (has bool, err error) {
	has, err = engine.Get(user)
	if err != nil {
		return false, err
	}
	return has, nil
}
