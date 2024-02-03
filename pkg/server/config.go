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
	ServiceVersion string
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

// VersionedPath builds the path string (version + path)
func (cfg *Config) VersionedPath(path string) string {
	return "/data/" + cfg.ServiceVersion + path
}

// SchemaVersionedEndpoint builds the schema endpoint string (schema + host + port + version + path)
func (cfg *Config) SchemaVersionedEndpoint(path string) string {
	if cfg.Port == 80 {
		return fmt.Sprintf("http://%s/data/%s%s", cfg.Host, cfg.ServiceVersion, path)
	}
	return fmt.Sprintf("http://%s:%d/data/%s%s", cfg.Host, cfg.Port, cfg.ServiceVersion, path)
}
