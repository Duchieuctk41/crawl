package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Item struct {
		Title        string
		Vendor       string
		Sale         string
		Price        string
		InitialPrice string
		TotalColor   string
		Images       []string
		Colors       []string
		Sizes        []string
		Sku          string
		Details      []string
	}

	Link struct {
		Link string
	}

	Collection struct {
		ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
		Name   string             `json:"name" bson:"name"`
		Page   int `json:"page" bson:"page"`
		Level1 []Level1
	}

	Level1 struct {
		Link string `json:"link" bson:"link"`
		Name string `json:"name" bson:"name"`
	}
)
