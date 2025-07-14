package main

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReverseTextResponse struct {
	Text         string `json:"text"`
	IsPalindrome bool   `json:"isPalindrome"`
}

func main() {
	router := gin.Default()
	router.GET("/", GetHandler)
	router.POST("/", FormHandler)
	router.PUT("/", PutHandler)
	router.Run("localhost:8080")
}

func getStatusCode(isPalindrome bool) int {
	if isPalindrome {
		return http.StatusAccepted
	}
	return http.StatusOK
}

func GetHandler(c *gin.Context) {
	text := c.Query("text") // параметр из самой ссылки локалхоста
	a, isPalindrome := ReverseString(text)
	txtColor := "green"
	if isPalindrome {
		txtColor = "red"
	}
	a = fmt.Sprintf(`<h1 style="color: %s">%s</h1>`, txtColor, a)

	c.String(getStatusCode(isPalindrome), a)

}

func FormHandler(c *gin.Context) {
	text := c.PostForm("text")
	a, isPalindrome := ReverseString(text)
	txtColor := "blue"
	if isPalindrome {
		txtColor = "red"
	}
	a = fmt.Sprintf(`<h1 style="color: %s">%s</h1>`, txtColor, a)

	c.String(getStatusCode(isPalindrome), a)
}

func PutHandler(c *gin.Context) {
	var body ReverseTextResponse
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ivalid body"})
		return
	}
	reversed, isPalindrome := ReverseString(body.Text)

	response := ReverseTextResponse{
		Text:         reversed,
		IsPalindrome: isPalindrome,
	}
	c.JSON(getStatusCode(isPalindrome), response)
}

func ReverseString(s string) (string, bool) {
	uppercase := strings.ToUpper(s) //нормализация строки в верхний регистр
	runes := []rune(uppercase)
	slices.Reverse(runes)
	reversed := string(runes)
	return reversed, uppercase == reversed
}
