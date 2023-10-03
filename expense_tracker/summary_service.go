package expense_tracker

import "github.com/sanjith9525/awesomeProject1/expense_tracker/models"

type SummaryService interface {
	GetUserSummary(userID int) (models.UserSummary, error)
	GetGroupSummary(groupID int) (models.GroupSummary, error)
}

type SummaryServiceImpl struct {
	groupRepo GroupRepo
}

func NewSummaryService(groupRepo GroupRepo) SummaryService {
	return &SummaryServiceImpl{
		groupRepo: groupRepo,
	}
}

func (s *SummaryServiceImpl) GetGroupSummary(groupID int) (models.GroupSummary, error) {
	groupInfo, err := s.groupRepo.FetchGroupInfo(groupID)
	if err != nil {
		return models.GroupSummary{}, err
	}
	groupSummary := s.getGroupSummaryFromInfo(groupInfo)
	return groupSummary, nil

}

func (s *SummaryServiceImpl) getGroupSummaryFromInfo(groupInfo Group) models.GroupSummary {
	owesTo := make([]models.Owes, 0)
	owedBy := make([]models.Owes, 0)
	var totalExpense float64
	for userFrom, userToMap := range groupInfo.users {
		for userTo, amount := range userToMap {
			if amount > 0 {
				owesTo = append(owesTo, models.Owes{
					FromUser: userFrom,
					ToUser:   userTo,
					Amount:   amount,
				})
				totalExpense += amount
			} else if amount < 0 {
				owedBy = append(owedBy, models.Owes{
					ToUser:   userTo,
					FromUser: userFrom,
					Amount:   amount,
				})
			}
		}
	}

	return models.GroupSummary{
		TotalExpense: totalExpense,
		OwesTo:       owesTo,
		Owed:         owedBy,
	}
}

func (s *SummaryServiceImpl) GetUserSummary(userID int) (models.UserSummary, error) {
	return models.UserSummary{}, nil
}
