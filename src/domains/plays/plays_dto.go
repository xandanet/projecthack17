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
