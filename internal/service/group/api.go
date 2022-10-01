package group

type GroupList struct {
}

func GetGroupList() {
	ImAppGroupGathers.GroupsMap.Range(func(key, value interface{}) bool {
		return true
	})
}
