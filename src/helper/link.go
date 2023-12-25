package helper

import (
	"context"
	"nxt/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetLinkByShortcode(shortcode string) (*models.Link, error) {
    var link models.Link
    err := db.Collection("links").FindOne(context.TODO(), bson.M{"short_code": shortcode}).Decode(&link)
    if err != nil {
        return nil, err
    }
    return &link, nil
}

func DeleteLinkByShortcode(shortcode string) error {
    _, err := db.Collection("links").DeleteOne(context.TODO(), bson.M{"short_code": shortcode})
    return err
}

func InsertLink(link *models.Link) error {
    _, err := db.Collection("links").InsertOne(context.TODO(), bson.M{
        "short_code": link.ShortCode,
        "url": link.URL,
        "pass_code": link.PassCode,
        "created_at": link.CreatedAt,
        "ip": link.IP,
    })
    return err
}

func TotalLinks() (int64, error) {
    count, err := db.Collection("links").CountDocuments(context.TODO(), bson.M{})
    if err != nil {
        return 0, err
    }
    return count, nil
}
