package main

import (
    "golang-crud-krs/config"
    "golang-crud-krs/models"
    "golang-crud-krs/routes"
)

func main() {
    config.ConnectDB()
    models.Migrate(config.DB)

    r := routes.SetupRouter()
    r.Run(":8080")
}
