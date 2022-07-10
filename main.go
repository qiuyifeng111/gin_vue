package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null;unique"`
	Password  string `gorm:"size:255"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机必须为11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return

		}
		if len(name) == 0 {
			name = RandomString(10)

		}
		log.Println(name, telephone, password)

		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
			return
		}
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		ctx.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run()
}
func RandomString(n int) string {
	var letters = []byte("asdfghjklQWERTYUIOP")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]

	}
	return string(result)
}
func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "gin_vue"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	return user.ID != 0
}
