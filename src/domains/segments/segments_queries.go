package segments

const (
	queryList = `SELECT id, start, end, content, pod_id, speaker, sentiment FROM subtitle;`

	queryListTextOnly = `SELECT content FROM subtitle LIMIT 100;`

	querySearchByText = `SELECT id, start, end, pod_id, content, speaker, sentiment FROM subtitle WHERE content = ?;`

	querySearchByNaturalSearchText = `SELECT id, start, end, pod_id, content, speaker, sentiment,
    	MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE) AS similarity FROM subtitle
		WHERE MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE);`
)
