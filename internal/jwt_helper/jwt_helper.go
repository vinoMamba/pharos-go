package jwt_helper

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user_id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
	})
	key, err := getHMACKey()
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

func getHMACKey() ([]byte, error) {

	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	keyPath := homePath + "/.pharos"

	bytes, err := ioutil.ReadFile(keyPath + "/hmac.key")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return bytes, nil
}

func GenerateHMACKey() ([]byte, error) {
	key := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
