package main

import (
	"github.com/heyyyyyyi/go-biz-admin/database"
	"github.com/heyyyyyyi/go-biz-admin/routes"
)

func main() {
	//connect to database (->)
	database.Connect()
	//initialize router (<-)
	r := routes.SetupRouter()
	r.Run(":8080")
}
