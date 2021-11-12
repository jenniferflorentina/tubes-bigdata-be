package router

import (
	controller "TubesBigData/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	// Add a todo
	app.Post("/todo", controller.AddTodo)

	// Get a todo
	app.Get("/todo/:id", controller.GetOne)
	app.Get("/todo", controller.GetAll)
	// Update a todo
	app.Put("/todo/:id", controller.UpdateOne)

	// Delete a todo
	app.Delete("/todo/:id", controller.DeleteOne)
	app.Delete("/todo", controller.DeleteMultiple)
}
