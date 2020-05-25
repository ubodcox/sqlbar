package rating

type (
	// WebUser struct
	WebUser struct {
		Name   string `json:"name"`
		Link   string `json:"link"`
		Rating int    `json:"rating"`
	}

	// WebRating struct
	WebRating struct {
		Users []WebUser `json:"users"`
	}
)

// NewWebRating func
func NewWebRating() *WebRating {
	return &WebRating{
		Users: make([]WebUser, 0),
	}
}
