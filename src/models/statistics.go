package models

import "time"

type Statistics struct {
	LinkId      string            `bson:"link_id" json:"link_id"`
	Date        string            `bson:"date" json:"date"`	
	Clicks      int               `bson:"clicks" json:"clicks"`
	CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
}
