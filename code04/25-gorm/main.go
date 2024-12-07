package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct { //User users
	ID      int    `gorm:"primarykey"`
	Keyword string `gorm:"column:keywords"`
	City    string
}

func (User) TableName() string {
	return "user"
}

// func read(client *gorm.DB, city string) *User {
// 	var user User
// 	// err := client.Select("id,city").Where("city = ?", city).Limit(1).First(&user).Error

// 	user.ID = 853985
// 	err := client.First(&user).Error
// 	checkError(err)

// 	return &user
// }

func main() {
	dataSourceName := "tester:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
	checkError(err)

	// user := read(client, "shanghai")
	// if user != nil {
	// 	fmt.Println(user)
	// } else {
	// 	fmt.Println("Not found")
	// }

	user := User{
		ID:      85397,
		Keyword: "golang",
		City:    "shanghai",
	}

	client.Create(&user)

	client.Model(User{}).Where("id = ?", 8538).Update("city", "beijing")
	client.Model(User{}).Where("id = ?", 8538).Updates(map[string]interface{}{"city": "beijing", "keywords": "golang"})
	client.Where("id = ?", 8538).Delete(&User{})

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
