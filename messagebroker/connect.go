package messagebroker

import (
    "context"
    "github.com/go-redis/redis/v9"
)

func NewConnect(ctx context.Context, password string, addr []string) (redis.UniversalClient, error) {
    client := redis.NewUniversalClient(&redis.UniversalOptions{
        Addrs:    addr,
        Password: password,
    })

    if err := client.Ping(ctx).Err(); err != nil {
        return nil, err
    }

    return client, nil
}
