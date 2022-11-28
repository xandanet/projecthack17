package subtitles

import "github.com/gobuffalo/nulls"

type SubtitleDTO struct {
	ID         int64        `db:"id"`
	Start      int64        `db:"start"`
	End        int64        `db:"end"`
	Content    string       `db:"content"`
	PodcastID  int64        `db:"pod_id"`
	Speaker    nulls.String `db:"speaker"`
	Sentiment  nulls.String `db:"sentiment"`
	Similarity float64      `db:"similarity"`
}

type SubtitleSearchInput struct {
	Text string `json:"text" form:"text" validate:"required"`
}

type TextSearchAnalysis struct {
	Text       string
	Similarity float64
}