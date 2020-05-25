package rating

type (
	// DBUser struct
	DBUser struct {
		Name   string
		Rating int
	}

	// DBRating struct
	DBRating struct {
		Users []DBUser
	}
)

// MakeWeb func
func (dbRating *DBRating) MakeWeb() (webRating *WebRating) {
	webRating = NewWebRating()
	for _, dbUser := range dbRating.Users {
		webUser := dbUser.MakeWeb()
		webRating.Users = append(webRating.Users, *webUser)
	}
	return
}

// MakeWeb func
func (dbUser *DBUser) MakeWeb() *WebUser {
	return &WebUser{
		Name:   dbUser.Name,
		Link:   "",
		Rating: dbUser.Rating,
	}
}
