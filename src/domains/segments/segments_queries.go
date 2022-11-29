package segments

const (
	queryList = `SELECT id, start, end, content, pod_id, speaker, sentiment FROM subtitle;`

	queryListTextOnly = `SELECT content FROM subtitle LIMIT 100;`

	querySearchByText = `SELECT id, start, end, pod_id, content, speaker, sentiment FROM subtitle WHERE content = ?;`

	querySearchByNaturalSearchText = `SELECT id, start, end, pod_id, content, speaker, sentiment,
    	MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE) AS similarity FROM subtitle
		WHERE MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE);`

	queryGetByID = `SELECT id, start, end, pod_id, content, speaker, sentiment FROM subtitle WHERE id = ?`

	queryGetSearchBySubtitleByID = `SELECT id, subtitle_id,search_id,click_count,first_clicked,last_clicked FROM search_subtitle WHERE id = ?`

	querySearchBySubtitleIdSearchId = `SELECT subtitle_id,search_id FROM search_subtitle WHERE subtitle_id = ? AND search_id=?`

	queryCreateSearchSubtitle = `INSERT INTO search_subtitle (subtitle_id,search_id) VALUES (:subtitle_id,:search_id)`

	queryUpdateCount = `UPDATE search_subtitle SET click_count = click_count+1 WHERE subtitle_id = ? AND search_id = ?`
)
