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
