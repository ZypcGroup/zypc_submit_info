package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

const (
	DataType = "mysql"
	Database = "infomation"
	Username = "root"
	Password = "axiu"
	Host     = "127.0.0.1"
	Port     = "3306"
)

var (
	datatype string
	database string
	username string
	password string
	host     string
	port     string
)

var engine *xorm.Engine

type User struct {
	UserId   int64 `xorm:"index"`
	UserName int
	Password string
	Time     time.Time `xorm:"index"`
	Flag     int
}

type Topic struct {
	UserId  int64 `xorm:"index"`
	Content string
	Flag    int
	Time    time.Time `xorm:"index"`
}

type Infomation struct {
	UserId int64 `xorm:"index"`
	Flag   int64
	Info1  string
	Info2  string
	Info3  string
	Info4  string
	Info5  string
	Info6  string
	Info7  string
	Info8  string
	Info9  string
	Info10 string
	Info11 string
	Info12 string
	Info13 string
	Info14 string
	Info15 string
	Info16 string
	Info17 string
	Info18 string
	Info19 string
	Info20 string
	Time   time.Time `xorm:"index"`
}

func init() {
	datatype = DataType
	database = Database
	host = Host
	port = Port
	username = Username
	password = Password
}

func RegisterDB() (err error) {

	engine, err = xorm.NewEngine(datatype, username+":"+password+"@/"+database+"?charset=utf8")
	fmt.Println(engine.Ping())
	if err != nil {
		return err
	}

	if ok, _ := engine.IsTableExist("user"); !ok {
		engine.CreateTables(new(User))
	}

	if ok, _ := engine.IsTableExist("topic"); !ok {
		engine.CreateTables(new(Topic))
	}

	if ok, _ := engine.IsTableExist("infomation"); !ok {
		engine.CreateTables(new(Infomation))
	}
	return nil
}

func AddUser(user *User) {

}

func ModifyUser() {

}

func AddInfomation(info *Infomation) {
	fmt.Println(info)
}

func ModifyInfomation() {

}

func SearchInfomation() {

}

func ShowAllInfomation() {

}

func ExportInfomation() {

}
