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