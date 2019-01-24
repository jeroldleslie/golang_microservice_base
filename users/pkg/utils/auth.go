package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"crypto/rsa"
	"os"
	"errors"
	"io/ioutil"
	error1 "go-microservice-base/users/pkg/errors"
	)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() (*JWTAuthenticationBackend, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}
	publicKey, err := getPublicKey()
	if err != nil {
		return nil, err
	}

	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: privateKey,
			PublicKey:  publicKey,
		}
	}
	return authBackendInstance, nil
}

func (backend *JWTAuthenticationBackend) GenerateToken(userid string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(tokenDuration)).Unix(),
		"iat": time.Now().Unix(),
		"sub": userid,
	}

	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyBytes, err := ioutil.ReadFile(os.Getenv("FIVEKILOMETER_PRIVATE_KEY"))
	if err != nil {
		return nil, errors.New(error1.CannotFindPrivateKey)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	publicKeyBytes, err := ioutil.ReadFile(os.Getenv("FIVEKILOMETER_PUBLIC_KEY"))
	if err != nil {
		return nil, errors.New(error1.CannotFindPublicKey)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}