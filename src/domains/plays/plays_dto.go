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
