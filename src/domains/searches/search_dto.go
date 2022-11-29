package searches

type SearchDTO struct {
	ID          int64  `db:"id"`
	Text        string `db:"text"`
	SearchCount int64  `db:"search_count"`
	LastUpdated string `db:"last_updated"`
}

type SearchInput struct {
	Text      string `db:"text"`
	Sentiment string `db:"sentiment"`
}
