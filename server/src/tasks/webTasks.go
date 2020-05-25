package tasks

type (
	// WebTask struct
	WebTask struct {
		GUID  string   `json:"guid"`
		Text  string   `json:"text"`
		Cols  []string `json:"cols"`
		Stars int      `json:"stars"`
	}

	// WebRating struct
	WebTasks struct {
		List []WebTask `json:"list"`
	}
)
