package helper

import (
	"context"
	"nxt/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func addClickToLink(link *models.Link, date string) {
    var statistics models.Statistics
    err := db.Collection("statistics").FindOne(context.TODO(), bson.M{"link_id": link.ShortCode, "date": date}).Decode(&statistics)
    if err != nil {
        statistics = models.Statistics{
            LinkId: link.ShortCode,
            Date: date,
            Clicks: 1,
        }
        _, err = db.Collection("statistics").InsertOne(context.TODO(), statistics)
        if err != nil {
        }
    } else {
        statistics.Clicks++
        _, err = db.Collection("statistics").UpdateOne(context.TODO(), bson.M{"link_id": link.ShortCode, "date": date}, bson.M{"$set": bson.M{"clicks": statistics.Clicks}})
        if err != nil {
        }
    }
}

func Clicks(link *models.Link) int {
	var statistics models.Statistics
	err := db.Collection("statistics").FindOne(context.TODO(), bson.M{"link_id": link.ShortCode}).Decode(&statistics)
	if err != nil {
		return 0
	}
	return statistics.Clicks
}

func ClickToLink(link *models.Link) error {
    addClickToLink(link, time.Now().Format("2006-01-02"))
    return nil
}
