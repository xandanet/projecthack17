package plays

const (
	queryCreate = `INSERT INTO plays(podcast_id, position, created_at) VALUES(?, ?, ?);`

	queryStatistics = `SELECT COUNT(*) AS total, position FROM plays
		WHERE podcast_id = ? AND DATE(created_at) BETWEEN ? AND ? 
		GROUP BY position;`

	queryPerDay = `SELECT COUNT(*) AS total, DATE(created_at) AS date FROM plays WHERE podcast_id = ? GROUP BY DATE(created_at) ORDER BY DATE ASC`

	querySegmentPopularity = `SELECT id as segment_id, start, end,
		(SELECT COUNT(*) FROM plays WHERE position BETWEEN segment.start AND segment.end AND plays.podcast_id = segment.pod_id) AS plays
		FROM segment WHERE pod_id = ?;`
)
