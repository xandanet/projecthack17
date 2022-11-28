package keywords

const (
	queryCreate = "INSERT INTO keywords(word) VALUE(?);"

	queryFindID = "SELECT id FROM keywords WHERE word = ?;"

	queryCreateRelationshipSubtitle = "INSERT INTO keyword_subtitle(keyword_id, subtitle_id) VALUES(?, ?);"
)
