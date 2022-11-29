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

type TopSearchesOutput struct {
	ID        int64  `db:"id"`
	Start     int64  `db:"start"`
	End       int64  `db:"end"`
	Content   string `db:"content"`
	PodcastID int64  `db:"pod_id"`
	Total     int64  `db:"total"`
	Podcast   string `db:"podcast"`
}

type SearchLocationsOutput struct {
	City     string `db:"city" json:"city"`
	Country  string `db:"country" json:"country"`
	Searches int64  `db:"searches" json:"searches"`
}
