package models

import "time"

type Statistics struct {
	Link        *Link             `bson:"link" json:"link"`
	Date        string            `bson:"date" json:"date"`
	Clicks      int               `bson:"clicks" json:"clicks"`
	CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
}
