// @title Public Library API
// @version 1.0
// @description A simple REST API to manage library books
// @host localhost:8080
// @BasePath /
// @schemes http

package main

import (
    "library-app/config"
    "library-app/src/routes"
)

func main() {
    config.Connect()
    r := routes.SetupRouter()
    r.Run(":8080")
}