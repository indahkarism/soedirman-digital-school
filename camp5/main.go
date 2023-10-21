package main

import (
	"tugas-sds/camp5/database"
	"tugas-sds/camp5/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	//Kumpulan Route Route

	app.Post("/", route.InsertData)
	app.Get("/", route.GetAllData)
	app.Get("/:id_user", route.GetUserByid)

	app.Delete("/:id_user", route.Delete)
	app.Put("/:id_user", route.Update)

	app.Listen(":3000")
}
