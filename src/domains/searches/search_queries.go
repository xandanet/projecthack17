package searches

const (
	queryList = `SELECT id, text, search_count, last_searched FROM search;`

	queryCreate = `INSERT INTO search (text, sentiment,has_result) VALUES (:text, :sentiment,1)`

	queryCreateNoResult = `INSERT INTO search (text, sentiment) VALUES (:text, :sentiment)`

	queryUpdateCount = `UPDATE search SET search_count = search_count+1 WHERE text = ?`

	queryFind = `SELECT id,text,search_count FROM search WHERE text = ?`

	queryByID = `SELECT id,text,search_count FROM search WHERE id = ?`

	/*queryTopSegmentFromSearch = `SELECT SUM(click_count) AS total, segment.id, start, end, content, pod_id, pods.title AS podcast
	FROM search_segment
	LEFT JOIN segment ON segment.id = segment_id
	LEFT JOIN pods ON pods.id = pod_id
	GROUP BY segment_id ORDER BY total DESC LIMIT ?`*/
	queryTopSegmentFromSearch = `SELECT SUM(search_count) AS total, text FROM search  WHERE has_result=1 GROUP BY text ORDER BY total DESC LIMIT 20;`

	queryTopSegmentNoResultFromSearch = `SELECT SUM(search_count) AS total, text FROM search WHERE has_result=0 GROUP BY text ORDER BY total DESC LIMIT 20;`

	queryGetSearchLocations = `SELECT city, country, count(*) AS searches FROM search_log GROUP BY country ORDER BY searches DESC LIMIT 10`
)
