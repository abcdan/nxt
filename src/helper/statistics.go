package helper

import (
	"context"
	"nxt/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func addClickToLink(link *models.Link, date string) {
    var statistics models.Statistics
    err := db.Collection("statistics").FindOne(context.TODO(), bson.M{"link": link, "date": date}).Decode(&statistics)
    if err != nil {
        statistics = models.Statistics{
            Link: link,
            Date: date,
            Clicks: 1,
        }
        _, err = db.Collection("statistics").InsertOne(context.TODO(), statistics)
        if err != nil {
        }
    } else {
        statistics.Clicks++
        _, err = db.Collection("statistics").UpdateOne(context.TODO(), bson.M{"link": link, "date": date}, bson.M{"$set": bson.M{"clicks": statistics.Clicks}})
        if err != nil {
        }
    }
}
func ClickToLink(link *models.Link) error {
    addClickToLink(link, time.Now().Format("2006-01-02"))
    return nil
}
