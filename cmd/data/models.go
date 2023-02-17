package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var client *redis.Client

func New(c *redis.Client) Models {
	client = c
	return Models{
		ImageData: ImageData{},
	}
}

type Models struct {
	ImageData ImageData
}

type ImageData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func (m *ImageData) insert(image ImageData) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Set(ctx, image.ID, image, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (m *ImageData) get(uuid string) (*ImageData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := client.Get(ctx, uuid).Result()

	if err == redis.Nil {
		return &ImageData{}, errors.New("not found")
	} else if err != nil {
		return &ImageData{}, errors.New("internal server error")
	}
	fmt.Println(value)
	return &ImageData{}, nil
}
