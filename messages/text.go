package messages

const (
	// InvalidContestTime : We have a contest in this range. ID: %v
	InvalidContestTime = "We have a contest in this range. ID: %v"

	// GeneralDBError : Error while working with database: %v
	GeneralDBError = "Error while working with database: %v"

	// MappingError : error in map %s to %s: %v
	MappingError = "error in map %s to %s: %v"

	// MetaHasContest : a contest already added to this meta. contestID: %v
	MetaHasContest = "a contest already added to this meta. contestID: %v"

	// ObjectNotFound : %s not found with %s:%v
	ObjectNotFound = "%s not found with %s:%v"

	// WrongIndexInSetAnsewr : User select index: %v and the lenght of questions is %v for question index %v in contestID %v
	WrongIndexInSetAnsewr = "User select index: %v and the lenght of questions is %v for question index %v in contestID %v"

	// ErrorInSplitCorrectAnswerIndices : Error while trying to split correct answer indices for question: %v => error : %v
	ErrorInSplitCorrectAnswerIndices = "Error while trying to split correct answer indices for question: %v => error : %v"
)
