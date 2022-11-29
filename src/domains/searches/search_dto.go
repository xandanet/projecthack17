package searches

type SearchDTO struct {
	ID           int64  `db:"id" json:"id"`
	Text         string `db:"text"  json:"text"`
	SearchCount  int64  `db:"search_count"  json:"search_count"`
	LastSearched string `db:"last_searched"  json:"last_searched"`
}

type SearchInput struct {
	Text      string `db:"text"`
	Sentiment string `db:"sentiment"`
}
