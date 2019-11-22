package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" bdinding:"required"`
	Age      int    `form:"age" json:"age"`}
var account map[string]User = make(map[string]User)
func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var user User
		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		username := user.Username
		if _, ok := account[username]; ok {
			message := "用户名" + username + "已存在"
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"message": message,
			})
		} else {
			account[username] = user
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"message": "注册成功",
			})
		}
		if v, ok := account[user.Username]; ok && v.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{
				"username":   v.Username,
				"password":     v.Password,
				"age":        v.Age,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message" : "账号或者密码有误",
			})
		}
	})
	router.Run()
}
