package server

import (
	"fmt"

	"github.com/mensatt/backend/pkg/utils"
)

// Config defines the configuration for the server
type Config struct {
	Host           string
	Port           int32
	MaxConnections int32
	DebugEnabled   bool
	JWT            utils.JWTKeyStoreConfig
	ImageAPIURL    string
	ImageAPIKey    string
}

// ListenEndpoint builds the endpoint string (host + port)
func (cfg *Config) ListenEndpoint() string {
	if cfg.Port == 80 {
		return cfg.Host
	}
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

// SchemaVersionedEndpoint builds the schema endpoint string (schema + host + port + version + path)
func (cfg *Config) SchemaVersionedEndpoint(path string) string {
	if cfg.Port == 80 {
		return fmt.Sprintf("http://%s/data/%s", cfg.Host, path)
	}
	return fmt.Sprintf("http://%s:%d/data/%s", cfg.Host, cfg.Port, path)
}
