package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}

func main() {
	db, err := gorm.Open("mysql", "root:1999@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("无法连接数据库")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	// 清空表数据
	db.Delete(&User{})

	// 创建多个记录
	users := []User{
		{Name: "Zhangsan", Email: "zhangsan@example.com", Age: 30},
		{Name: "Lisi", Email: "lisi@example.com", Age: 38},
		{Name: "Wangwu", Email: "wangwu@example.com", Age: 32},
	}

	for _, user := range users {
		db.Create(&user)
		fmt.Printf("Created user: %v\n", user)
	}

	// 查询记录
	var foundUser User
	db.First(&foundUser, "name = ?", "Zhangsan")
	fmt.Printf("Found user: %v\n", foundUser)

	// 更新记录
	db.Model(&foundUser).Update("Name", "Zhangdasan")
	fmt.Printf("Updated user: %v\n", foundUser)

	// 删除记录
	db.Delete(&foundUser)
	fmt.Printf("Deleted user: %v\n", foundUser)

	// 输出表信息
	var usersList []User
	db.Find(&usersList)
	fmt.Println("All users:")
	for _, u := range usersList {
		fmt.Printf("%v\n", u)
	}
}
