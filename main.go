package main

import (
	"log"

	api "budget-service/handlers"
)

func main() {
	log.Println("budget_app_started_up")
	api.Start()
	// fmt.Println(models.NewBudgetRepo("us-west-2", "").Get("349755"))
}
