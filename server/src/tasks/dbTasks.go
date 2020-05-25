package tasks

type (
	// WebTask struct
	DbTask struct {
		GUID  string
		Text  string
		Cols  []string
		Stars int
	}

	// WebRating struct
	DbTasks struct {
		List []WebTask
	}
)

/*// MakeWeb func
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
*/
