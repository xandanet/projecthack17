package keywords

type KeywordDTO struct {
	ID         int64  `db:"id"`
	SubtitleID int64  `db:"subtitle_id"`
	Keyword    string `db:"keyword"`
	Label      string `db:"label"`
}
