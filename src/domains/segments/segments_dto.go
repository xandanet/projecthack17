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
	Plays      int64        `db:"plays"`
}

type SegmentListInput struct {
	PodcastID int64 `json:"podcast_id" form:"podcast_id" validate:"required"`
}
type SegmentSearchInput struct {
	Text string `json:"text" form:"text" validate:"required"`
}

type TextSearchAnalysis struct {
	Text       string
	Similarity float64
}

type SearchSegmentDTO struct {
	SearchID   int64
	SegmentDTO []SegmentDTO
}

type SearchSegmentInput struct {
	SegmentId int64  `json:"segment_id" db:"segment_id" form:"segment_id" validate:"required"`
	SearchId  int64  `json:"search_id" db:"search_id" form:"search_id" validate:"required"`
	IpAddress string `json:"ip_address" db:"ip_address" form:"ip_address" validate:"required"`
}

type SearchSegmentOutput struct {
	ID           int64  `db:"id"`
	SegmentId    int64  `db:"segment_id"`
	SearchId     int64  `db:"search_id"`
	ClickCount   int64  `db:"click_count"`
	FirstClicked string `db:"first_clicked"`
	LastClicked  string `db:"last_clicked"`
}

type SearchLogInput struct {
	SearchId  int64  ` db:"search_id"`
	IpAddress string ` db:"ip_address"`
	Region    string ` db:"region"`
	City      string ` db:"city"`
	Country   string ` db:"country"`
}
