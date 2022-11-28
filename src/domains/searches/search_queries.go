package searches

const (
	queryList = `SELECT id, text, count, last_updated FROM search;`

	queryCreate = `INSERT INTO search (text) VALUES (:text)`

	queryUpdateCount = `UPDATE search SET search_count = search_count+1 WHERE text = ?`

	queryFind = `SELECT id,text,count FROM search WHERE text = ?`
)
