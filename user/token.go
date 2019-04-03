package user

import (
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var privateKey = []byte("MHcCAQEEIHgfULIcCnEGuL1f1lNhP9dhIupvsYv1OVJjIc3J3jrFoAoGCCqGSM49AwEHoUQDQgAETFSC75YsHVIS02G7Z1t0Au81F0J5ljlD6e9JnTQ3eY2VxYBcs8DwGrQp6VlcYjKXF4Eiy2oy8nbRfsiZZjisow==")

var privateKeyString = hex.EncodeToString(privateKey)

func createToken(m *User) (string, error) {
	token := jwt.New(jwt.SigningMethodRSA)

	token.Claims["email"] = m.Email
	token.Claims["password"] = m.Password
	token.Claims["expire"] = time.Now().Add(time.Minute * 10).Unix()
	tokenString, err := token.SignedString(privateKey)
	return tokenString, err
}

func checkToken(token *jwt.Token) error {
	token, err := jwt.Parse(privateKeyString, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err == nil && token.Valid {
		return nil
	}
	if err != nil {
		return err
	}
}
