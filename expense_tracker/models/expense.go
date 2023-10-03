package models

type Expense struct {
	UserSplit   []UserSplit `json:"user_split"`
	ExpenseType string      `json:"expense_type"`
	Amount      float64     `json:"amount"`
}

type UserSplit struct {
	UserID int     `json:"user_id"`
	Paid   float64 `json:"paid"`
	Val    float64 `json:"val"`
}

type SplitValues struct {
	UserSplit map[int]UserOwed
}

type UserOwed struct {
	UserId int
	Amount float64
}

type ExpenseRequest struct {
	GroupID int     `json:"group_id"`
	Expense Expense `json:"expense"`
}
