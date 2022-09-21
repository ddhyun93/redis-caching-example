package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
)

type Client struct {
	client *redis.Client
}

func NewRedisClient() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:637912",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	ping := client.Ping(ctx)
	if ping.String() != "ping: PONG" {
		return nil, errors.New(ping.String())
	}

	return &Client{client: client}, nil
}

func (c *Client) Get(key string) ([]byte, error) {
	return nil, nil
}

func (c *Client) Set(key string) ([]byte, error) {
	return nil, nil
}
