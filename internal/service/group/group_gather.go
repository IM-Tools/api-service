package group

import "sync"

var (
	ImAppGroupGathers = AppGroupGathers{
		GroupsMap: sync.Map{},
	}
)

type AppGroupGathers struct {
	GroupsMap sync.Map
}

func (gathers *AppGroupGathers) SetGroups(groups *Groups) {
	gathers.GroupsMap.Store(groups.Id, groups)
}
