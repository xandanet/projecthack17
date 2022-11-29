package plays

const (
	queryCreate = `INSERT INTO plays(podcast_id, position, created_at) VALUES(?, ?, ?);`

	queryStatistics = `SELECT COUNT(*) AS total, position FROM plays
		WHERE podcast_id = ? AND DATE(created_at) BETWEEN ? AND ? 
		GROUP BY position;`
)
