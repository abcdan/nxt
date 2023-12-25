package models

import (
	"time"
)

type Link struct {
	IP         string             `bson:"ip" json:"ip"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	ShortCode  string             `bson:"short_code" json:"short_code"`
	URL        string             `bson:"url" json:"url"`
	PassCode   *string            `bson:"passcode,omitempty" json:"passcode,omitempty"`
}
