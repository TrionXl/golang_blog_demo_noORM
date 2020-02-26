package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go_blog/service"
)

const servingAddress = ":1323"
const (
	authname = "admin"
	authpw   = "secret"
)

func main() {
	e := echo.New()
	service.Router(e)                                                                          //启动路由
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) { //简单认证
		if username == authname && password == authpw {
			return true, nil
		}
		return false, nil
	}))
	e.Logger.Fatal(e.Start(servingAddress))
}
