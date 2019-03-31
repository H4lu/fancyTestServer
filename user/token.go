package user

import (
	"github.com/dgrijalva/jwt-go"
)

const privateKey string = "MHcCAQEEIHgfULIcCnEGuL1f1lNhP9dhIupvsYv1OVJjIc3J3jrFoAoGCCqGSM49
AwEHoUQDQgAETFSC75YsHVIS02G7Z1t0Au81F0J5ljlD6e9JnTQ3eY2VxYBcs8Dw
GrQp6VlcYjKXF4Eiy2oy8nbRfsiZZjisow=="

func createToken(m *User) (string, error) {
	token := jwt.New(jwt.SigningMethodECDSA)

	token.Claims["email"] = m.Email
	token.Claims["password"] = m.Password
	token.Claims["expire"] = time.Now().Add(time.Minute * 10).Unix()
}