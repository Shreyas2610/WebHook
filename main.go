package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Route to create a new post
	app.Post("/posts", createPost)

	// Start the server on port 3000
	app.Listen(":8080")
}

func createPost(c *fiber.Ctx) error {
	// Parse JSON request body into a map interface
	var results map[string]interface{}

	jsonbody := c.Body()

	json.Unmarshal([]byte(jsonbody), &results)

	fmt.Println("***********************************************************")
	fmt.Println("WebHook Payload Data")
	fmt.Println(string(jsonbody))
	fmt.Println("***********************************************************")

	// Return the same JSON data that was accepted
	return c.Status(fiber.StatusCreated).JSON(results)
}
