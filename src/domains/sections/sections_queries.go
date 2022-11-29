package sections

const (
	queryList = `SELECT section.id, section.start, section.end, section.content, content_pt, section.speaker, sentiment
       FROM section
       LEFT JOIN segment ON subtitle_id = segment.id
       WHERE pod_id = ?;`
)
