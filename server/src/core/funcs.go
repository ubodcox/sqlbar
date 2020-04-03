package core

import (
	"fmt"
	"net/http"
	"time"

	//"strings"
	"crypto/rand"
	"math"
	"sqlbar/server/src/logs"
)

const (
	// MaxShortDescrLen const
	MaxShortDescrLen = 48
	// MaxDescrLen const
	MaxDescrLen = 98
)

const (
	// StatusNew const
	StatusNew = "NEW"
	// StatusInWork const
	StatusInWork = "IN_WORK"
	// StatusDone const
	StatusDone = "DONE"
	// StatusCanceled const
	StatusCanceled = "CANCELED"
)

// TimeToA func
func TimeToA(time time.Time) string {
	return time.Format("02-01-2006 15:04:05")
}

// EnableCors func
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// TrimToLen func
func TrimToLen(s string, leng int) string {
	if len([]rune(s)) > leng {
		bytes := len(s)
		runes := len([]rune(s))
		//log.Println("SS", bytes, runes, leng)
		index := math.Floor(float64(bytes / runes * leng))
		//log.Println("SSS", index, bytes, runes)
		return s[:int(index)] + "..."
	}

	return s
}

// MakeGUID func
func MakeGUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logs.Log.Error(err)
	}
	uuid := fmt.Sprintf("%x%x%x%x%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	//fmt.Println(uuid)
	return uuid
}
