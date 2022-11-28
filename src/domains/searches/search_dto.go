package searches

type SearchDTO struct {
	ID          int64  `db:"id"`
	Text        int64  `db:"text"`
	SearchCount int64  `db:"search_count"`
	LastUpdated string `db:"last_updated"`
}

type SearchInput struct {
	Text string ` db:"text"`
}
