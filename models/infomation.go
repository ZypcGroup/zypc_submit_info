package models

import (
	"fmt"
	// "github.com/go-xorm/xorm"
	"time"
)

func AddInfomation(info *Infomation) bool {

	fmt.Println(info)
	connectDB()
	info.Time = time.Now()
	_, err := engine.Insert(info)
	if err == nil {
		return true
	}
	return false

}

func ModifyInfomation() {

}

func SearchInfomation() {

}

func ShowAllInfomation() {

}

func ExportInfomation() {

}
