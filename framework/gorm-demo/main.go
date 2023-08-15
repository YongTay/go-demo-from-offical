package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	EmployeeId string `gorm:"primarykey"`
	FirstName  string
	LastName   string
}

func (e Employee) String() string {
	// 反射获取所有的值和类型，用于输出
	aValue := reflect.ValueOf(e)
	aType := reflect.TypeOf(e)
	n := aValue.NumField()
	// 用于存储键：值
	kv := make([]string, n)
	for i := 0; i < n; i++ {
		fv := aValue.Field(i)
		f := aType.Field(i)
		kv[i] = fmt.Sprintf("%s:%s", f.Name, fv)
	}
	return "{ " + strings.Join(kv, ", ") + " } "
}

func check(err error) {
	if nil != err {
		panic(err)
	}
}

func main() {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	url := os.Getenv("MYSQL_URL")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/myemployees?charset=utf8mb4&parseTime=True&loc=Local", user, password, url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	check(err)
	conns, err := db.DB()
	check(err)

	conns.SetMaxIdleConns(10)
	conns.SetMaxOpenConns(30)
	conns.SetConnMaxLifetime(time.Hour)

	err = conns.Ping()
	check(err)
	defer conns.Close()

	var data []Employee
	// db.First(&data) // 查询一行数据
	// db.First(&data, "employee_id=?", "121") // 条件查询
	// db.Model(Employee{EmployeeId: "121"}).First(&data) // 条件查询
	// db.Where("employee_id > 120 and employee_id < 130").Find(&data) // 条件查询
	// result := db.Find(&data) // 查询所有数据
	// db.Select("FirstName").Limit(10).Find(&data) // 查询指定字段
	// fmt.Println(result.RowsAffected)
	fmt.Println(data)
}
