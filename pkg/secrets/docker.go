package secrets

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"os"
	"strings"
)

const defaultSecretsDir = "/run/secrets"

// DockerSecrets defines a structure that shall contain our Docker secrets
type DockerSecrets struct {
	secretsDir string
	secrets    map[string]string
}

// CreateSecretsContainer creates an instance of DockerSecrets
// return os.ErrNotExist if secrets dir not exists
func CreateSecretsContainer() (*DockerSecrets, error) {
	dockerSecrets := &DockerSecrets{secretsDir: defaultSecretsDir, secrets: map[string]string{}}
	err := dockerSecrets.loadAll()
	return dockerSecrets, err
}

// Get returns one secret by secretName
func (ds *DockerSecrets) Get(secretName string) (string, error) {
	if _, ok := ds.secrets[secretName]; !ok {
		return "", fmt.Errorf("secret doesn't exsist: %s", secretName)
	}
	return ds.secrets[secretName], nil
}

// Reads all secrets on the specified path in the secretsDir
func (ds *DockerSecrets) loadAll() error {
	secretsDir := ds.secretsDir
	err := isDir(secretsDir)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(secretsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := ds.read(file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// Reads a secret from file
func (ds *DockerSecrets) read(file string) error {
	buf, err := ioutil.ReadFile(ds.secretsDir + "/" + file)
	if err != nil {
		return err
	}
	ds.secrets[file] = strings.TrimSpace(string(buf))
	return nil
}

// Checks if the given path is a directory
func isDir(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.Mode().IsDir() {
		return fmt.Errorf("is not a directory: %s", path)
	}
	return nil
}

// Unmarshal unmarshals the secrets into a Struct
func (ds *DockerSecrets) Unmarshal(output interface{}) error {
	return decode(ds.secrets, defaultDecoderConfig(output))
}

// defaultDecoderConfig returns default mapsstructure.DecoderConfig
func defaultDecoderConfig(output interface{}) *mapstructure.DecoderConfig {
	return &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
}

// A wrapper around mapstructure.Decode
func decode(input interface{}, config *mapstructure.DecoderConfig) error {
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
