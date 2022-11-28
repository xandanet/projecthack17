package keywords

type KeywordDTO struct {
	ID   int64  `db:"id"`
	Word string `db:"word"`
}
