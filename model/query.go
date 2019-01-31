package model

func saveMetaQuery() string {
	return `
	INSERT INTO contest (appID, title, prize, beginTime, duration, itemDuration, neededCorrectors, allowedCorrectorUsageTimes, allowCorrectTilQuestion, neededTickets)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`
}
