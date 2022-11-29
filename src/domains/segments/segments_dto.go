package segments

import "github.com/gobuffalo/nulls"

type SegmentDTO struct {
	ID         int64        `db:"id"`
	Start      int64        `db:"start"`
	End        int64        `db:"end"`
	Content    string       `db:"content"`
	PodcastID  int64        `db:"pod_id"`
	Speaker    nulls.String `db:"speaker"`
	Sentiment  nulls.String `db:"sentiment"`
	Similarity float64      `db:"similarity"`
}

type SegmentSearchInput struct {
	Text string `json:"text" form:"text" validate:"required"`
}

type TextSearchAnalysis struct {
	Text       string
	Similarity float64
}

type SearchSubtitleDTO struct {
	SearchID    int64
	SubtitleDTO []SegmentDTO
}

type SearchSubtitleInput struct {
	SubtitleId int64 `json:"subtitleId" form:"subtitle_id" validate:"required"`
	SearchId   int64 `json:"searchId" form:"search_id" validate:"required"`
}

type SearchSubtitleOutput struct {
	ID           int64 `db:"id"`
	SubtitleId   int64 `db:"subtitle_id"`
	SearchId     int64 `db:"search_id"`
	ClickCount   int64 `db:"click_count"`
	FirstClicked int64 `db:"first_clicked"`
	LastClicked  int64 `db:"last_clicked"`
}
