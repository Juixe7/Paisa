package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	RDB *redis.Client
}

func Connect(ctx context.Context, url string) (*Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("parsing redis URL: %w", err)
	}

	rdb := redis.NewClient(opts)

	// Verify connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		rdb.Close()
		return nil, fmt.Errorf("pinging redis: %w", err)
	}

	return &Client{RDB: rdb}, nil
}

func (c *Client) Close() error {
	if c.RDB != nil {
		return c.RDB.Close()
	}
	return nil
}
