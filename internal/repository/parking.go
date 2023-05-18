package repository

import (
    "app/internal/model"
    "context"
    "github.com/redis/go-redis/v9"
    "strconv"
    "time"
)

type RedisRepository struct {
    rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) *RedisRepository {
    return &RedisRepository{rdb: rdb}
}

func (p *RedisRepository) SaveLocally(ctx context.Context, parking model.TaxiParking) {
    globalID := strconv.Itoa(parking[0].GlobalId)
    p.rdb.Set(ctx, globalID, parking, time.Hour*3)
}

func (p *RedisRepository) DeleteLocal(ctx context.Context, globalID int) {
    p.rdb.Del(ctx, strconv.Itoa(globalID))
}

func (p *RedisRepository) CheckExpiration(ctx context.Context, globalID int) time.Duration {
    duration := p.rdb.TTL(ctx, strconv.Itoa(globalID))
    return duration.Val()
}
