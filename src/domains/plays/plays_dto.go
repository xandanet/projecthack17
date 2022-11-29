package plays

type PlayDTO struct {
	ID        int64  `db:"id"`
	PodcastID int64  `db:"podcast_id"`
	Position  int64  `db:"position"`
	CreatedAt string `db:"created_at"`
}

type PlayCreateInput struct {
	PodcastID int64  `db:"podcast_id" json:"podcast_id" form:"podcast_id" validate:"required"`
	Position  int64  `db:"position" json:"position" form:"position" validate:"required"`
	CreatedAt string `db:"created_at"`
}

type PlayStatisticsInput struct {
	PodcastID int64  `json:"podcast_id" form:"podcast_id" validate:"required"`
	StartDate string `json:"start_date" form:"start_date" validate:"required"`
	EndDate   string `json:"end_date" form:"end_date" validate:"required"`
}

type PlayStatisticsOutput struct {
	Count     int64 `db:"total" json:"count"`
	Timestamp int64 `db:"position" json:"timestamp"`
}

type PlayStatisticsPerDayInput struct {
	PodcastID int64 `json:"podcast_id" form:"podcast_id" validate:"required"`
}

type PlayStatisticsPerDayOutput struct {
	Count int64  `db:"total" json:"count"`
	Date  string `db:"date" json:"date"`
}

type PlaySegmentPopularityOutput struct {
	SegmentID int64 `db:"segment_id" json:"segment_id"`
	Start     int64 `db:"start" json:"start"`
	End       int64 `db:"end" json:"end"`
	Plays     int64 `db:"plays" json:"plays"`
}
