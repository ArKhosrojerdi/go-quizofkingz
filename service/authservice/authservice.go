package authservice

import "time"

type Config struct {
	SignKey                   string
	AccessExpirationDuration  time.Duration
	RefreshExpirationDuration time.Duration
	AccessSubject             string
	RefreshSubject            string
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}
