package main

import (
	"github.com/sanjith9525/awesomeProject1/expense_tracker"
	"github.com/sanjith9525/awesomeProject1/handlers"
	"net/http"
)

func main() {

	expressHandler := initApp()

	http.HandleFunc("/api/v1/expenses/", expressHandler.ExpensesHandler)
	http.HandleFunc("/api/v1/groups/{id}/summary/", expressHandler.GroupSummaryHandler)
	//http.HandleFunc("/api/v1/users/{id}/summar", IndividulSummaryHandler)

	http.ListenAndServe(":8080", nil)
}

func initApp() handlers.ExpenseHandler {
	groupRepo := expense_tracker.NewGroupRepo()
	expenseService := expense_tracker.NewExpenseService(groupRepo)
	summaryService := expense_tracker.NewSummaryService(groupRepo)
	expenseHandler := handlers.ExpenseHandler{
		ExpenseService: expenseService,
		SummaryService: summaryService,
	}
	return expenseHandler

}
