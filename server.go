package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", getUsers)
	e.Logger.Fatal(e.Start(":4000"))
}

func getUsers(c echo.Context) error {
	a := []User{
		{Name: "リンクバルタロウ", Email: "test1@test.com"},
		{Name: "リンクバルジロウ", Email: "test2@test.com"}}
	return c.JSON(http.StatusOK, a)
}
