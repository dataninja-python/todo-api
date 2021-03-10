package controllers

import (
	"github.com/Kamva/mgm"
	"github.com/dataninja-python/todo-api/models"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// route to get all todos
func GetAllTodos(ctx *fiber.Ctx) {
	// execute finding all todos
	collection := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}

	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
}
