package main

import (
	"fmt"
	"slices"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", GetHandler)
	router.POST("/", FormHandler)
	router.PUT("/", PutHandler)
	router.Run("localhost:8080")
}

func GetHandler(c *gin.Context) {
	text := c.Query("text") // параметр из самой ссылки локалхоста
	a, isPalindrome := ReverseString(text)
	txtColor := "green"
	if isPalindrome {
		txtColor = "red"
	}
	a = fmt.Sprintf(`<h1 style="color: %s">%s</h1>`, txtColor, a)
	c.String(http.StatusOK, a) // вернутое значение

}

func FormHandler(c *gin.Context) {
	text := c.PostForm("text")
	a, isPalindrome := ReverseString(text)
	txtColor := "blue"
	if isPalindrome {
		txtColor = "red"
	}
	a = fmt.Sprintf(`<h1 style="color: %s">%s</h1>`, txtColor, a)
	c.String(http.StatusOK, a)
}

func PutHandler(c *gin.Context) {

}

func ReverseString(s string) (string, bool) {
	runes := []rune(s)
	slices.Reverse(runes)
	reversed := string(runes)
	return reversed, s == reversed
}
