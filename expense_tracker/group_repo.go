package expense_tracker

type GroupRepo interface {
	AddExpenseToGroup(groupID int, userFrom int, userTo int, amount float64) error
	FetchUserGroups(userID int) ([]int, error)
	FetchGroupInfo(groupID int) (Group, error)
}

type GroupRepoImpl struct {
	groups     map[int]Group
	userGroups map[int][]int
}

type Group struct {
	users map[int]map[int]float64
}

func NewGroupRepo() GroupRepo {
	return &GroupRepoImpl{
		groups:     make(map[int]Group),
		userGroups: make(map[int][]int),
	}
}

func (g GroupRepoImpl) AddExpenseToGroup(groupID int, userFrom int, userTo int, amount float64) error {
	if _, ok := g.groups[groupID]; !ok {
		g.groups[groupID] = Group{
			users: make(map[int]map[int]float64),
		}
	}
	if _, ok := g.groups[groupID].users[userFrom]; !ok {
		g.groups[groupID].users[userFrom] = make(map[int]float64)
		g.userGroups[userFrom] = append(g.userGroups[userFrom], groupID)
	}
	if _, ok := g.groups[groupID].users[userTo]; !ok {
		g.groups[groupID].users[userTo] = make(map[int]float64)
		g.userGroups[userFrom] = append(g.userGroups[userFrom], groupID)
	}

	g.groups[groupID].users[userFrom][userTo] += amount
	g.groups[groupID].users[userTo][userFrom] -= amount

	return nil

}

func (g GroupRepoImpl) FetchUserGroups(userID int) ([]int, error) {
	return g.userGroups[userID], nil
}

func (g GroupRepoImpl) FetchGroupInfo(groupID int) (Group, error) {
	return g.groups[groupID], nil
}
