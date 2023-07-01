package config

import "time"

const (
	DefaultGracefulShutdownTimeOut = 30 * time.Second

	DefaultDatabaseMaxIdleConns       = 3
	DefaultDatabaseMaxOpenConns       = 5
	DefaultDatabaseConnMaxLifetime    = 1 * time.Hour
	DefaultDatabasePingInterval       = 1 * time.Second
	DefaultDatabaseRetryAttempts      = 3
	DefaultDatabaseReconnectFactor    = 2
	DefaultDatabaseReconnectMinJitter = 100 * time.Millisecond
	DefaultDatabaseReconnectMaxJitter = 1 * time.Second

	DefaultRedisDialTimeout  = 5 * time.Second
	DefaultRedisWriteTimeout = 2 * time.Second
	DefaultRedisReadTimeout  = 2 * time.Second
	DefaultRedisCacheTTL     = 15 * time.Minute

	DefaultJetstreamMaxPending = 256
	DefaultJetstreamMaxAge     = 24 * time.Hour

	DefaultAsynqConcurrency = 10
	DefaultAsynqRetry       = 3
	DefaultAsynqRetention   = 15 * time.Minute
)
