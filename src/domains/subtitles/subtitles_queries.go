package subtitles

const (
	queryList = `SELECT id, start, end, content, pod_id, speaker, sentiment FROM subtitle;`

	queryListByPodcast = `SELECT id, start, end, content, pod_id, speaker, sentiment FROM subtitle WHERE pod_id = ?;`

	queryListAllText = `SELECT content FROM subtitle;`

	queryListTextOnly = `SELECT content FROM subtitle WHERE pod_id = ?;`

	querySearchByText = `SELECT id, start, end, pod_id, content, speaker, sentiment FROM subtitle WHERE content = ?;`

	querySearchByNaturalSearchText = `SELECT id, start, end, pod_id, content, speaker, sentiment,
    	MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE) AS similarity FROM subtitle
		WHERE MATCH (content) AGAINST (? IN NATURAL LANGUAGE MODE);`
)
