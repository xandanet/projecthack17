package podcasts

const (
	queryMaxDuration = `SELECT MAX(END) AS duration, pod_id FROM segment GROUP BY pod_id`
)
