package tasks

type (
	//ResultList struct
	ResultList struct {
		IsCorrect bool                     `json:"iscorrect"`
		List      []map[string]interface{} `json:"list"`
	}
)
