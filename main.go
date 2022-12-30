package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	config := gorm.Config{
		Dialector: mysql.New(mysql.Config{
			DSN: "root:rootpassword@tcp(localhost:3308)/go_admin?charset=utf8&parseTime=True&loc=Local"}),
	}

	db, err := gorm.Open(config)
	if err != nil {
		panic("Database is down!")
	}
	fmt.Println(db)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8000"))
}
