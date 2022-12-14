package segments

const (
	queryList = `SELECT id, start, end, content, pod_id, speaker, sentiment,
       (SELECT COUNT(*) FROM plays WHERE podcast_id = segment.pod_id AND position BETWEEN segment.start AND segment.end) AS plays
       FROM segment WHERE pod_id = ?;`

	queryListTextOnly = `SELECT content FROM segment;`

	querySearchByText = `SELECT id, start, end, pod_id, content, speaker, sentiment, pods.title AS podcast
		FROM segment
		LEFT JOIN pods ON (pods.id = segment.pod_id)
		WHERE content = ?;`

	querySearchByNaturalSearchText = `SELECT segment.id, start, end, pod_id, content, speaker, sentiment, pods.title AS podcast,
    	MATCH (content) AGAINST (? IN BOOLEAN MODE) + MATCH (pods.title, pods.description) AGAINST (? IN BOOLEAN MODE) AS similarity
		FROM segment
    	LEFT JOIN pods ON (pods.id = segment.pod_id)
		WHERE MATCH (content) AGAINST (? IN BOOLEAN MODE) + MATCH (pods.title, pods.description) AGAINST (? IN BOOLEAN MODE)
		ORDER BY similarity DESC;`

	queryGetByID = `SELECT id, start, end, pod_id, content, speaker, sentiment FROM segment WHERE id = ?`

	queryGetSearchBySubtitleByID = `SELECT id, segment_id,search_id,click_count,first_clicked,last_clicked FROM search_segment WHERE id = ?`

	querySearchBySubtitleIdSearchId = `SELECT id, segment_id,search_id,click_count,first_clicked,last_clicked FROM search_segment WHERE segment_id = ? AND search_id=?`

	queryCreateSearchSubtitle = `INSERT INTO search_segment (segment_id,search_id) VALUES (:segment_id,:search_id)`

	queryUpdateCount = `UPDATE search_segment SET click_count = click_count+1 WHERE segment_id = ? AND search_id = ?`

	queryCreateLog = `INSERT INTO search_log (search_id,ip,region,city,country) VALUES (:search_id,:ip_address,:region,:city,:country)`
)
