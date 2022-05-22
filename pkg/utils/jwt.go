package utils

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/golang-jwt/jwt/v4"
)

//JWTKeyStoreConfig defines the options for JWTKeyStore
type JWTKeyStoreConfig struct {
	PrivateKeyPath string
	PublicKeyPath  string
	Algorithm      string
}

type JWTKeyStore struct {
	config      *JWTKeyStoreConfig
	algorithm   jwt.SigningMethod
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
	pubKeyBytes []byte
}

func LoadJWTKeys(jwtConfig *JWTKeyStoreConfig) (*JWTKeyStore, error) {
	pubBytes, err := ioutil.ReadFile(jwtConfig.PublicKeyPath)
	if err != nil {
		return nil, err
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, err
	}

	privBytes, err := ioutil.ReadFile(jwtConfig.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return nil, err
	}
	privKey.Precompute()

	algo := jwt.GetSigningMethod(jwtConfig.Algorithm)
	if algo == nil {
		return nil, fmt.Errorf("unexpected signing method: %v", jwtConfig.Algorithm)
	}

	return &JWTKeyStore{
		config:      jwtConfig,
		algorithm:   algo,
		privateKey:  privKey,
		publicKey:   pubKey,
		pubKeyBytes: pubBytes,
	}, nil
}

// GenerateJWT generates JWT with the provided data
func (ks *JWTKeyStore) GenerateJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(ks.algorithm, claims)
	tokenString, err := token.SignedString(ks.privateKey)
	return tokenString, err
}

// ParseJWT parses the JWT
func (ks *JWTKeyStore) ParseJWT(tokenString string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if ks.algorithm != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ks.pubKeyBytes, nil
	})
	return err
}
