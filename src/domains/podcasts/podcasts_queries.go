package podcasts

const (
	queryMaxDuration = `SELECT MAX(END) AS duration, pod_id FROM segment GROUP BY pod_id`

	queryList = `SELECT id, title, description, file_name, season, episode, DATE(streamed_on) AS streamed_on, file_name_pt IS NOT NULL AS has_subtitles  FROM pods;`

	queryInterventions = `SELECT SUM(end-start) AS total_spoken, COUNT(*) AS interventions, speaker,       
		(SELECT COUNT(*) FROM plays WHERE podcast_id = segment.pod_id AND position BETWEEN segment.start AND segment.end) AS plays
		FROM segment WHERE pod_id = ? GROUP BY speaker`

	querySentiment = `SELECT start, sentiment FROM segment WHERE pod_id = ?;`
)
