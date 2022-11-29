package podcasts

type PodcastDuration struct {
	ID       int64 `db:"pod_id"`
	Duration int64 `db:"duration"`
}
