package config

type (
	// InfoFile struct
	InfoFile struct {
		InfoConfig `json:"config"`
	}

	// InfoConfig struct
	InfoConfig struct {
		InfoServer `json:"server"`
	}

	// InfoServer struct
	InfoServer struct {
		Version  string `json:"version"`
	}

	//TODO: add change log
)

