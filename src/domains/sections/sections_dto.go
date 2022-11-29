package sections

import "github.com/gobuffalo/nulls"

type SectionDTO struct {
	ID        int64        `db:"id"`
	Start     int64        `db:"start"`
	End       int64        `db:"end"`
	Content   string       `db:"content"`
	ContentPT string       `db:"content_pt"`
	Speaker   nulls.String `db:"speaker"`
	Sentiment string       `db:"sentiment"`
}

type SectionListInput struct {
	PodcastID int64 `json:"podcast_id" form:"podcast_id" validate:"required"`
}
