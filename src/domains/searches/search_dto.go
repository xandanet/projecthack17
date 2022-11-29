package searches

type SearchDTO struct {
	ID           int64  `db:"id"`
	Text         string `db:"text"`
	SearchCount  int64  `db:"search_count"`
	LastSearched string `db:"last_searched"`
}

type SearchInput struct {
	Text      string `db:"text"`
	Sentiment string `db:"sentiment"`
}

type TopSearchesOutput struct {
	ID        int64  `db:"id"`
	Start     int64  `db:"start"`
	End       int64  `db:"end"`
	Content   string `db:"content"`
	PodcastID int64  `db:"pod_id"`
	Total     int64  `db:"total"`
	Podcast   string `db:"podcast"`
}
