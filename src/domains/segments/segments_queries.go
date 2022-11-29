package segments

const (
	queryList = `SELECT id, start, end, content, pod_id, speaker, sentiment FROM segment;`

	queryListTextOnly = `SELECT content FROM segment LIMIT 100;`

	querySearchByText = `SELECT id, start, end, pod_id, content, speaker, sentiment FROM segment WHERE content = ?;`

	querySearchByNaturalSearchText = `SELECT id, start, end, pod_id, content, speaker, sentiment,
    	MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE) AS similarity FROM segment
		WHERE MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE);`

	queryGetByID = `SELECT id, start, end, pod_id, content, speaker, sentiment FROM segment WHERE id = ?`

	queryGetSearchBySubtitleByID = `SELECT id, segment_id,search_id,click_count,first_clicked,last_clicked FROM search_segment WHERE id = ?`

	querySearchBySubtitleIdSearchId = `SELECT id, segment_id,search_id,click_count,first_clicked,last_clicked FROM search_segment WHERE segment_id = ? AND search_id=?`

	queryCreateSearchSubtitle = `INSERT INTO search_segment (segment_id,search_id) VALUES (:segment_id,:search_id)`

	queryUpdateCount = `UPDATE search_segment SET click_count = click_count+1 WHERE segment_id = ? AND search_id = ?`
)
