package main

import "gin-bug-example/api"

func main() {
	app := api.GetApp()
	app.Run("localhost:8080")
}
