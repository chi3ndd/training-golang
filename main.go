package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Request struct {
	A int `query:"A"`
	B int `query:"B"`
}

func main() {
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
	// Start webservice
	e.Start("0.0.0.0:8910")
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
