package main

import (
	"fmt"
	"slices"

	"net/http"

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

func GetHandler(c *gin.Context) {
	text := c.Query("text") // параметр из самой ссылки локалхоста
	a, isPalindrome := ReverseString(text)
	txtColor := "green"
	if isPalindrome {
		txtColor = "red"
	}
	a = fmt.Sprintf(`<h1 style="color: %s">%s</h1>`, txtColor, a)

	// Возвращаем 304 для палиндромов, иначе 200
	if isPalindrome {
		c.String(http.StatusNotModified, a)
	} else {
		c.String(http.StatusOK, a)
	}

}

func FormHandler(c *gin.Context) {
	text := c.PostForm("text")
	a, isPalindrome := ReverseString(text)
	txtColor := "blue"
	if isPalindrome {
		txtColor = "red"
	}
	a = fmt.Sprintf(`<h1 style="color: %s">%s</h1>`, txtColor, a)

	// Возвращаем 304 для палиндромов, иначе 200
	if isPalindrome {
		c.String(http.StatusNotModified, a)
	} else {
		c.String(http.StatusOK, a)
	}
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
	if isPalindrome {
		c.JSON(http.StatusNotModified, response)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func ReverseString(s string) (string, bool) {
	runes := []rune(s)
	slices.Reverse(runes)
	reversed := string(runes)
	return reversed, s == reversed
}
