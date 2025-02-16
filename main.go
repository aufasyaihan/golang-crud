package main

import (
    "golang-crud-users/config"
    "golang-crud-users/models"
    "golang-crud-users/routes"
)

func main() {
    config.ConnectDB()
    models.Migrate(config.DB)

    r := routes.SetupRouter()
    r.Run(":8080")
}
