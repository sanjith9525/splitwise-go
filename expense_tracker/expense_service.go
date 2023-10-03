package expense_tracker

import (
	"fmt"
	"github.com/sanjith9525/awesomeProject1/expense_tracker/models"
)

type ExpenseAdder interface {
	AddExpenseToGroup(groupID int, expense models.Expense) error
}

type ExpenseEngine interface {
	GetExpenseSplit(expense models.Expense) ([]models.SplitValues, error)
}

type ExpenseService struct {
	groupRepo GroupRepo
}

func NewExpenseService(groupRepo GroupRepo) ExpenseAdder {
	return &ExpenseService{
		groupRepo: groupRepo,
	}
}

func (e *ExpenseService) AddExpenseToGroup(groupID int, expense models.Expense) error {
	expenseType := expense.ExpenseType
	expenseCalculationEngine := getExpenseEngine(expenseType)
	splitValues, err := expenseCalculationEngine.GetExpenseSplit(expense)
	if err != nil {
		return err
	}
	for _, splitValue := range splitValues {
		for userID, split := range splitValue.UserSplit {
			err = e.groupRepo.AddExpenseToGroup(groupID, userID, split.UserId, split.Amount)
			if err != nil {
				return fmt.Errorf("error while adding expense to group: %w", err)
			}
		}
	}
	return nil
}

func getExpenseEngine(expenseType string) ExpenseEngine {
	if expenseType == "equal" {
		return &EqualExpenseEngine{}
	}
	return nil
}

type EqualExpenseEngine struct {
}

func (e *EqualExpenseEngine) GetExpenseSplit(expense models.Expense) ([]models.SplitValues, error) {
	err := e.amountValidation(expense)
	if err != nil {
		return nil, err
	}
	splitValues := make([]models.SplitValues, 0)
	eachUserSplitAmount := expense.Amount / float64(len(expense.UserSplit))
	owedUsers := make([]models.UserOwed, 0)
	owingUsers := make([]models.UserOwed, 0)
	for _, userSplit := range expense.UserSplit {
		pendingAmount := userSplit.Paid - eachUserSplitAmount
		if pendingAmount > 0 {
			owedUsers = append(owedUsers, models.UserOwed{
				UserId: userSplit.UserID,
				Amount: pendingAmount,
			})
		} else {
			owingUsers = append(owingUsers, models.UserOwed{
				UserId: userSplit.UserID,
				Amount: -pendingAmount,
			})
		}
	}

	for _, owingUsers := range owingUsers {
		splitValue := models.SplitValues{
			UserSplit: make(map[int]models.UserOwed),
		}
		for _, owedUser := range owedUsers {
			if owingUsers.Amount > owedUser.Amount {
				splitValue.UserSplit[owedUser.UserId] = models.UserOwed{
					UserId: owedUser.UserId,
					Amount: owedUser.Amount,
				}
				owingUsers.Amount -= owedUser.Amount
				owedUser.Amount = 0
			} else {
				splitValue.UserSplit[owedUser.UserId] = models.UserOwed{
					UserId: owedUser.UserId,
					Amount: owingUsers.Amount,
				}
				owedUser.Amount -= owingUsers.Amount
				owingUsers.Amount = 0
			}
		}
		splitValues = append(splitValues, splitValue)
	}
	return splitValues, nil
}

func (e *EqualExpenseEngine) amountValidation(expense models.Expense) error {
	totalAmount := 0.0
	for _, userSplit := range expense.UserSplit {
		totalAmount += userSplit.Paid
	}
	if totalAmount != expense.Amount {
		return fmt.Errorf("amount mismatch")
	}
	return nil
}
