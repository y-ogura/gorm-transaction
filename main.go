package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/y-ogura/gorm-transaction/db"
)

func main() {
	e := echo.New()

	db.ConnectDB()
	db.Migrate()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "hello")
	})

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
