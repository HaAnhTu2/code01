package main

import (
	"os"

	"github.com/HaAnhTu2/code01/db"
	"github.com/HaAnhTu2/code01/handler"
	"github.com/HaAnhTu2/code01/log"
	"github.com/HaAnhTu2/code01/reponsitory/repo_impl"
	"github.com/labstack/echo/v4"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "tusha",
		Password: "postgres",
		DbName:   "01042024",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	UserHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	e.POST("/user/create", UserHandler.HandleCreateUser)
	e.Logger.Fatal(e.Start(":3000"))
}
