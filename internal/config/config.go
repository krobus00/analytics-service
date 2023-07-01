package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	serviceName    = ""
	serviceVersion = ""
)

func ServiceName() string {
	return serviceName
}

func ServiceVersion() string {
	return serviceVersion
}

func DurableID() string {
	return fmt.Sprintf("%s-durable", serviceName)
}

func QueueGroup() string {
	return fmt.Sprintf("%s-queue-group", serviceName)
}

func Env() string {
	return viper.GetString("env")
}

func LogLevel() string {
	return viper.GetString("log_level")
}

func PortHTTP() string {
	return viper.GetString("ports.http")
}

func GracefulShutdownTimeOut() time.Duration {
	cfg := viper.GetString("graceful_shutdown_timeout")
	return parseDuration(cfg, DefaultGracefulShutdownTimeOut)
}

func DatabaseDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		DatabaseUsername(),
		DatabasePassword(),
		DatabaseHost(),
		DatabaseName(),
		DatabaseSSLMode())
}

func DatabaseHost() string {
	return viper.GetString("database.host")
}

func DatabaseName() string {
	return viper.GetString("database.database")
}

func DatabaseUsername() string {
	return viper.GetString("database.username")
}

func DatabasePassword() string {
	return viper.GetString("database.password")
}

func DatabaseSSLMode() string {
	if viper.IsSet("database.sslmode") {
		return viper.GetString("database.sslmode")
	}
	return "disable"
}

func DatabasePingInterval() time.Duration {
	return parseDuration(viper.GetString("database.ping_interval"), DefaultDatabasePingInterval)
}

func DatabaseRetryAttempts() float64 {
	if viper.GetInt("database.retry_attempts") > 0 {
		return float64(viper.GetInt("database.retry_attempts"))
	}
	return DefaultDatabaseRetryAttempts
}

func DatabaseMaxIdleConns() int {
	if viper.GetInt("database.max_idle_conns") <= 0 {
		return DefaultDatabaseMaxIdleConns
	}
	return viper.GetInt("database.max_idle_conns")
}

func DatabaseMaxOpenConns() int {
	if viper.GetInt("database.max_open_conns") <= 0 {
		return DefaultDatabaseMaxOpenConns
	}
	return viper.GetInt("database.max_open_conns")
}

func DatabaseConnMaxLifetime() time.Duration {
	if !viper.IsSet("database.conn_max_lifetime") {
		return DefaultDatabaseConnMaxLifetime
	}
	return time.Duration(viper.GetInt("database.conn_max_lifetime")) * time.Millisecond
}

func DatabaseConnReconnectFactor() int {
	if viper.GetInt("database.conn_reconnect_factor") <= 0 {
		return DefaultDatabaseMaxOpenConns
	}
	return viper.GetInt("database.conn_reconnect_factor")
}

func DatabaseConnReconnectMinJitter() time.Duration {
	cfg := viper.GetString("database.conn_reconnect_min_jitter")
	return parseDuration(cfg, DefaultDatabaseReconnectMinJitter)
}

func DatabaseConnReconnectMaxJitter() time.Duration {
	cfg := viper.GetString("database.conn_reconnect_Max_jitter")
	return parseDuration(cfg, DefaultDatabaseReconnectMaxJitter)
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config not found")
		}
		return err
	}
	return nil
}

func parseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}
