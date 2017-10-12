package strava

import (
	"time"
)

type PhotoSummary struct {
	Id             int64             `json:"id"`
	UniqueId       string            `json:"unique_id"`
	AthleteId      int64             `json:"athlete_id"`
	ActivityId     int64             `json:"activity_id"`
	ActivityName   string            `json:"activity_name"`
	ResourceState  int               `json:"resource_state"`
	Reference      string            `json:"ref"`
	UID            string            `json:"uid"`
	Caption        string            `json:"caption"`
	Type           string            `json:"type"`
	Source         int               `json:"source"`
	UploadedAt     time.Time         `json:"uploaded_at"`
	CreatedAt      time.Time         `json:"created_at"`
	CreatedAtLocal time.Time         `json:"created_at_local"`
	Urls           map[string]string `json:"urls"`
	Sizes          map[string][2]int `json:"sizes"`
	DefaultPhoto   bool              `json:"default_photo"`
	Location       Location          `json:"location"`
}
