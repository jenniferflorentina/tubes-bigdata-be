package router

import (
	controller "TubesBigData/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	// Add a todo
	app.Post("/add", controller.AddTodo)

	// Get a todo
	app.Get("/getOne/:id", controller.GetOne)
	app.Get("/getAll", controller.GetAll)
	// Update a todo
	app.Put("/updateOne/:id", controller.UpdateOne)

	// Delete a todo
	app.Delete("/deleteOne/:id", controller.DeleteOne)
	app.Delete("/deleteMultiple", controller.DeleteMultiple)
}
