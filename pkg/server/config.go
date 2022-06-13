package server

import (
	"fmt"

	"github.com/mensatt/backend/pkg/utils"
)

// ServerConfig defines the configuration for the server
type ServerConfig struct {
	Host           string
	Port           int32
	ServiceVersion string
	DebugEnabled   bool
	JWT            utils.JWTKeyStoreConfig
	AssetsDir      string
	ImageProcessor utils.ImageProcessorConfig
}

// ListenEndpoint builds the endpoint string (host + port)
func (cfg *ServerConfig) ListenEndpoint() string {
	if cfg.Port == 80 {
		return cfg.Host
	}
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

// VersionedEndpoint builds the endpoint string (host + port + version)
func (cfg *ServerConfig) VersionedPath(path string) string {
	return "/" + cfg.ServiceVersion + path
}

// SchemaVersionedEndpoint builds the schema endpoint string (schema + host + port + version)
func (cfg *ServerConfig) SchemaVersionedEndpoint(path string) string {
	if cfg.Port == 80 {
		return fmt.Sprintf("http://%s/%s%s", cfg.Host, cfg.ServiceVersion, path)
	}
	return fmt.Sprintf("http://%s:%d/%s%s", cfg.Host, cfg.Port, cfg.ServiceVersion, path)
}
