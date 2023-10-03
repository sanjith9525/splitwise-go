package models

type GroupSummary struct {
	TotalExpense float64 `json:"total_expense"`
	OwesTo       []Owes  `json:"owes_to"`
	Owed         []Owes  `json:"owed"`
}

type Owes struct {
	FromUser int     `json:"from_user"`
	ToUser   int     `json:"to_user"`
	Amount   float64 `json:"amount"`
}

type UserSummary struct {
	TotalExpense   float64          `json:"total_expense"`
	GroupWise      []GroupWise      `json:"group_wise"`
	IndividualWise []IndividualWise `json:"individual_wise"`
}

type GroupWise struct {
	GroupID int     `json:"group_id"`
	Owes    float64 `json:"owes"`
	Owed    float64 `json:"owed"`
}

type IndividualWise struct {
	UserID int     `json:"user_id"`
	Owes   float64 `json:"owes"`
	Owed   float64 `json:"owed"`
}
type GroupSummaryRequest struct {
}
