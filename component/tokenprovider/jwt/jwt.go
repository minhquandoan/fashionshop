package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/minhquandoan/fashionshop/component/tokenprovider"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{
		secret: secret,
	}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (provider *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt: time.Now().UTC().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(provider.secret))
	if err != nil {
		return nil, tokenprovider.ErrEncodingToken
	}

	return &tokenprovider.Token{
		Token: myToken,
		Expiry: expiry,
		Created: time.Now(),
	}, nil
}

func (provider *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(provider.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return &claims.Payload, nil
}

