package podcasts

type PodcastDTO struct {
	ID           int64  `db:"id" json:"id"`
	Title        string `db:"title" json:"title"`
	Description  string `db:"description" json:"description"`
	Filename     string `db:"file_name" json:"filename"`
	Season       int64  `db:"season" json:"season"`
	Episode      int64  `db:"episode" json:"episode"`
	StreamedOn   string `db:"streamed_on" json:"streamed_on"`
	HasSubtitles bool   `db:"has_subtitles" json:"has_subtitles"`
}

type PodcastDuration struct {
	ID       int64 `db:"pod_id" json:"pod_id"`
	Duration int64 `db:"duration" json:"duration"`
}

type PodcastInput struct {
	ID int64 `json:"id" form:"id" validate:"required"`
}

type PodcastInterventionsOutput struct {
	TimeSpoken    int64  `db:"total_spoken" json:"time_spoken"`
	Interventions int64  `db:"interventions" json:"interventions"`
	Plays         int64  `db:"plays" json:"plays"`
	Speaker       string `db:"speaker" json:"speaker"`
}

type PodcastSentimentOutput struct {
	Start     int64  `db:"start" json:"start"`
	Sentiment string `db:"sentiment" json:"sentiment"`
}

type BookmarkInput struct {
	PodI     int64  `db:"pod_id" json:"pod_id" form:"pod_id" validate:"required"`
	Position int64  `db:"position" json:"position" form:"position" validate:"required"`
	Notes    string `db:"notes" json:"notes" form:"notes"`
}

type BookmarkSearchInput struct {
	PodId int64 `json:"pod_id" form:"pod_id" validate:"required"`
}

type GetBookmarkSearchOutput struct {
	PodID    int64  `db:"pod_id" json:"pod_id" form:"pod_id" validate:"required"`
	Position int64  `db:"position" json:"position" form:"position" validate:"required"`
	Notes    string `db:"notes" json:"notes" form:"notes" validate:"required"`
}
