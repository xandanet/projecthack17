package searches

const (
	queryList = `SELECT id, text, search_count, last_updated FROM search;`

	queryCreate = `INSERT INTO search (text, sentiment) VALUES (:text, :sentiment)`

	queryUpdateCount = `UPDATE search SET search_count = search_count+1 WHERE text = ?`

	queryFind = `SELECT id,text,search_count FROM search WHERE text = ?`

	queryByID = `SELECT id,text,search_count FROM search WHERE id = ?`
)
