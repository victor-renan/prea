package security

import (
	"github.com/golang-jwt/jwt/v5"
	"prea/internal/common"
)

const (
	DefaultAlg string = "HS256"
	JwtEnv     string = "JWTKEY"
)

var SignKey []byte

func init() {
	SignKey = []byte(common.GetEnv(JwtEnv))
}

type Jwt[M any] struct{}

type DataClaims struct {
	Data any `json:"data"`
	*jwt.RegisteredClaims
}

func (j Jwt[M]) CreateToken(object M, claims jwt.RegisteredClaims) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(DefaultAlg))

	token.Claims = &DataClaims{
		object,
		&claims,
	}

	encoded, err := token.SignedString(SignKey)
	if err != nil {
		return "", err
	}

	return encoded, nil
}

func (j Jwt[M]) DecodeToken(token string) (DataClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	}

	parsed, err := jwt.ParseWithClaims(token, &DataClaims{}, keyFunc)
	if err != nil {
		return DataClaims{}, err
	}

	return *parsed.Claims.(*DataClaims), nil
}
