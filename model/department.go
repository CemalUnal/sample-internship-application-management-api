package model

// simulates enumeration for the Department info, and it is not persisted in the DB.
const(
	Marketing = "Marketing"
	Design = "Design"
	Development = "Development"
	CEO = "CEO"
)

func GetDepartmentsAsArray() []string {
	return []string { Marketing, Design, Development, CEO }
}
