package main

import (
	"fmt"
	"regexp"
	"slices"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ReverseTextResponse struct {
	Text         string `json:"text"`
	IsPalindrome bool   `json:"isPalindrome"`
}

var (
	textRegex = regexp.MustCompile(`^[а-яА-ЯёЁa-zA-Z0-9\s]+$`)
)

func main() {
	router := gin.Default()
	router.GET("/", WebHandler) // Объединённый обработчик для GET
	router.POST("/", WebHandler)
	router.PUT("/", PutHandler)
	router.Run("localhost:8080")
}

func validateText(text string) error {
	if len(text) == 0 {
		return fmt.Errorf("текст не может быть пустым")
	}
	if len(text) > 50 {
		return fmt.Errorf("текст слишком длинный (максимум 50 символов)")
	}
	if !textRegex.MatchString(text) {
		return fmt.Errorf("текст может содержать только буквы, цифры и пробелы")
	}
	return nil
}

func getStatusCode(isPalindrome bool) int {
	if isPalindrome {
		return http.StatusAccepted
	}
	return http.StatusOK
}

func WebHandler(c *gin.Context) {
	var text string
	if c.Request.Method == "GET" {
		text = c.Query("text")
	} else {
		text = c.PostForm("text")
	}

	if err := validateText(text); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	reversed, isPalindrome := ReverseString(text)

	// Цвет в зависимости от метода
	var txtColor string
	if c.Request.Method == "GET" {
		txtColor = "green"
	} else {
		txtColor = "blue"
	}
	if isPalindrome {
		txtColor = "red" // Палиндром всегда красный
	}

	htmlResponse := fmt.Sprintf(`<h1 style="color: %s">%s</h1>`, txtColor, reversed)
	c.String(getStatusCode(isPalindrome), htmlResponse)
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
	runes := []rune(s)
	slices.Reverse(runes)
	reversed := string(runes)
	return reversed, s == reversed
}
