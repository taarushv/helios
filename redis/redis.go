package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.TODO()

type RedisDB struct {
    Client *redis.Client
}

func NewRedisClient(address string) (*RedisDB, error) {
    client := redis.NewClient(&redis.Options{
        Addr: address,
        password: "",
        DB: 0,
    })

    if err := client.Ping(Ctx).Err(); err != nil {
        return nil, err
    }

    return &RedisDB{
        Client: client,
    }, nil

}

// Example Table for Uniswap:
// |   hash    |     from  |    to   |   fromAmt    |    toAmt  |
// |     0x..  |     0x..  |    0x.. |    5000      |    10000  |
// |     0x..  |     0x..  |    0x.. |    5000      |    10000  |
// |     0x..  |     0x..  |    0x.. |    5000      |    10000  |
// |     0x..  |     0x..  |    0x.. |   5000       |    10000  |
// sets relational-like data for a given txn
// example key: 'uniswap:txnHash' 
// example value: from 0x.. to 0x.. fromAmt 5000 toAmt 10000
func SetDataFor(table string, key string, value string) error {

}

