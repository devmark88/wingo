package q

import "fmt"

func getUserTopic(userID string) string {
	return fmt.Sprintf("user:%s", userID)
}
func getQuestionTopic(contestID uint) string {
	return fmt.Sprintf("contest%v", contestID)
}
func getDeadlineTopic(contestID uint) string {
	return fmt.Sprintf("contest%v", contestID)
}
