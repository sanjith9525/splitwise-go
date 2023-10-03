package handlers

import (
	"encoding/json"
	"github.com/sanjith9525/awesomeProject1/expense_tracker"
	"github.com/sanjith9525/awesomeProject1/expense_tracker/models"
	"net/http"
	"strconv"
)

type ExpenseHandler struct {
	ExpenseService expense_tracker.ExpenseAdder
	SummaryService expense_tracker.SummaryService
}

func (e ExpenseHandler) ExpensesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var expense models.ExpenseRequest
		err := json.NewDecoder(r.Body).Decode(&expense)
		if err != nil {
			ReturnWithError(w, err, http.StatusBadRequest)
		}

		err = e.ExpenseService.AddExpenseToGroup(expense.GroupID, expense.Expense)
		if err != nil {
			ReturnWithError(w, err, http.StatusInternalServerError)
		}
		ReturnWithJson(w, "Expense added successfully", http.StatusOK)
	}

}

func ReturnWithJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(statusCode)
}

func ReturnWithError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
	w.WriteHeader(statusCode)
}

func (e ExpenseHandler) GroupSummaryHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("group_id")
	id, err := strconv.Atoi(groupId)
	if err != nil {
		ReturnWithError(w, err, http.StatusBadRequest)
	}
	summary, err := e.SummaryService.GetGroupSummary(id)
	if err != nil {
		ReturnWithError(w, err, http.StatusInternalServerError)
	}
	ReturnWithJson(w, summary, http.StatusOK)
}
