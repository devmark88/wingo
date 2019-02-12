package messages

const (
	// InvalidContestTime : contest time is not valid
	InvalidContestTime = "We have a contest in this range. ID: %v"

	// GeneralDBError : global database error
	GeneralDBError = "Error while working with database: %v: %s"

	// MappingError : error while mapping objects
	MappingError = "error in map %s to %s: %v"

	// MetaHasContest : contest has question
	MetaHasContest = "a contest already added to this meta. contestID: %v"

	// ObjectNotFound : an object not found
	ObjectNotFound = "%s not found with %s:%v"
)
