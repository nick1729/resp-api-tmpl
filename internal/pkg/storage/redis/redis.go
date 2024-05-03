package redis

import (
	"context"
	"time"

	"github.com/nick1729/resp-api-tmpl/internal/pkg/config"
	"github.com/nick1729/resp-api-tmpl/internal/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	client *redis.Client
}

const createNewPoolTimeout = 5 * time.Second

var ErrKeyNotFound = errors.New("key not found")

func New(ctx context.Context, cfg config.Redis) (Service, error) {
	var resp Service

	ctx, cancel := context.WithTimeout(ctx, createNewPoolTimeout)
	defer cancel()

	options := &redis.Options{
		Addr:     cfg.BuildAddressString(),
		Password: cfg.Pass,
		DB:       cfg.DB,
	}

	client := redis.NewClient(options)

	err := client.Ping(ctx).Err()
	if err != nil {
		return resp, errors.Wrap(err, "checking redis connection")
	}

	resp.client = client

	return resp, nil
}

func (p Service) Close() {
	p.client.Close()
}

func (s Service) Ping(ctx context.Context) error {
	return s.client.Ping(ctx).Err()
}

func (s Service) Get(ctx context.Context, key string) (string, error) {
	resp, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrKeyNotFound
		}

		return "", errors.Wrap(err, "getting the value")
	}

	return resp, nil
}

func (s Service) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	_, err := s.client.Set(ctx, key, value, ttl).Result()
	if err != nil {
		return errors.Wrap(err, "setting the value")
	}

	return nil
}

func (s Service) SetNX(ctx context.Context, key, value string, ttl time.Duration) (bool, error) {
	resp, err := s.client.SetNX(ctx, key, value, ttl).Result()
	if err != nil {
		return false, errors.Wrap(err, "setting nx the value")
	}

	return resp, nil
}

func (s Service) Delete(ctx context.Context, key string) error {
	_, err := s.client.Del(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrKeyNotFound
		}

		return errors.Wrap(err, "deleting the value")
	}

	return nil
}

func (s Service) SRem(ctx context.Context, key string, members ...any) error {
	_, err := s.client.SRem(ctx, key, members...).Result()
	if err != nil {
		return errors.Wrap(err, "removing members from the set")
	}

	return nil
}

func (s Service) SMembers(ctx context.Context, key string) ([]string, error) {
	resp, err := s.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "getting members of the set")
	}

	return resp, nil
}

func (s Service) SAdd(ctx context.Context, key string, members ...any) error {
	_, err := s.client.SAdd(ctx, key, members...).Result()
	if err != nil {
		return errors.Wrap(err, "adding members to the set")
	}

	return nil
}

func (s Service) SIsMember(ctx context.Context, key string, member any) (bool, error) {
	resp, err := s.client.SIsMember(ctx, key, member).Result()
	if err != nil {
		return false, errors.Wrap(err, "checking key of the set")
	}

	return resp, nil
}
