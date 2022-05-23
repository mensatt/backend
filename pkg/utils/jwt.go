package utils

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

//JWTKeyStoreConfig defines the options for JWTKeyStore
type JWTKeyStoreConfig struct {
	PrivateKeyPath string
	PublicKeyPath  string
	Algorithm      string
	TimeoutSec     int32
}

type JWTKeyStore struct {
	timeout     time.Duration
	algorithm   jwt.SigningMethod
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
}

func InitJWTKeyStore(jwtConfig *JWTKeyStoreConfig) (*JWTKeyStore, error) {
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

	timeout := time.Duration(jwtConfig.TimeoutSec) * time.Second

	return &JWTKeyStore{
		timeout:     timeout,
		algorithm:   algo,
		privateKey:  privKey,
		publicKey:   pubKey,
	}, nil
}

type CustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT generates JWT with the provided data
func (ks *JWTKeyStore) GenerateJWT(userID uuid.UUID) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ks.timeout)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(ks.algorithm, claims)
	tokenString, err := token.SignedString(ks.privateKey)
	return tokenString, err
}

// ParseJWT parses the JWT
func (ks *JWTKeyStore) ParseJWT(tokenString string) (uuid.UUID, error) {
	var claims CustomClaims
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if ks.algorithm != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return ks.publicKey, nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	return claims.UserID, nil
}
