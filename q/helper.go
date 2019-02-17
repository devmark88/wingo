package q

import "fmt"

// GetUserTopicName : get topic name for user
func GetUserTopicName(userID string) string {
	return fmt.Sprintf("tpc:user:%s", userID)
}

// GetQuestionTopicName : get topic name for question
func GetQuestionTopicName(contestMetaID uint) string {
	return fmt.Sprintf("ccontest%v", contestMetaID)
}

// GetDeadlineTopicName : get topic name for deadline
func GetDeadlineTopicName(contestMetaID uint) string {
	return fmt.Sprintf("tpc:contest%v", contestMetaID)
}
