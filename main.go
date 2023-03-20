package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Request struct {
	A int `query:"A"`
	B int `query:"B"`
}

func main2() {
	// Init
	e := echo.New()
	// Config print log
	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}][${remote_ip} -> ${host}][${protocol}-${method}][${status}]: ${uri} | in: ${bytes_in} - out: ${bytes_out} | ${latency_human} | error: ${error}\n",
		Output: os.Stdout,
	})
	e.Use(logger)
	// Khai bao route den ham xu ly
	e.GET("/print", PrintABC)
	e.GET("/login", Login)
	e.GET("/test/:idid/test", Test)
	// Start webservice
	e.Start("0.0.0.0:8910")
}

type Request3 struct {
	ID int `param:"idid"`
}

func Test(c echo.Context) error {
	var r Request3
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(400, "loi loi")
	}
	fmt.Println(r.ID)

	return c.JSON(200, "OK")
}

func PrintABC(c echo.Context) error {
	// Doan code binding du lieu tu query vao 1 struct ma minh dinh nghia
	var r Request
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(500, "Loi roi")
	}
	// Return ket qua
	return c.JSON(200, r.A+r.B)
	// address := c.Request().RemoteAddr
	// if strings.Contains(address, "127.0.0.1") {
	// 	// Ban
	// 	return c.JSON(403, "May da bi chan")
	// } else {
	// 	// Allow
	// 	return c.JSON(200, "OK")
	// }
}

type Request2 struct {
	Username string `query:"username"`
	Password string `query:"password"`
}

func Login(c echo.Context) error {
	// Buoc 1: Binding vao struct & verify
	r, err := verifyLogin(c)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	// Xu ly
	if r.Username == "longlong" && r.Password == "chienchienchien" {
		return c.JSON(200, "Dang nhap thanh cong")
	}

	return c.JSON(401, "Dang nhap that bai")
}

func verifyLogin(c echo.Context) (Request2, error) {
	var r Request2
	err := c.Bind(&r)
	if err != nil {
		return Request2{}, err
	}
	if len(r.Username) < 6 {
		return Request2{}, errors.New("username qua ngan")
	}

	if len(r.Password) < 8 {
		return Request2{}, errors.New("password qua ngan")
	}

	return r, nil
}

// Bai tap: Viet 1 API dang nhap.
// Yeu cau: username, password truyen qua param (query)
// Neu do dai username < 6 => return code 400, noi dung username qua ngan
// Neu do dai password < 8 => return code 400, noi dung password qua ngan
// Neu username == longlong va password == chienchienchien thi tra ve code 200, noi dung dang nhap thanh cong
// Truong hop con lai thi tra ve code 401, dang nhap khong hop le
// Goi y
// 1. Khai bao 1 struct Request2 co thuoc tinh username, password va co tag de con binding vao nha
// 2. Viet ham xu ly luon o trong file main.go, co signature func ...(c echo.Context) error
// 3. Khai bao route e.GET("/login", *Ten cua function*)
