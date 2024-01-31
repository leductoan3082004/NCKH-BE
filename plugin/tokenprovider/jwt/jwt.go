package jwtProvider

import (
	"flag"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lequocbinh04/go-sdk/logger"
	"nckh-BE/plugin/tokenprovider"
	"time"
)

type jwtProvider struct {
	secretKey string
	prefix    string
	name      string
	logger    logger.Logger
}

func NewJwtProvider(name, prefix string) *jwtProvider {
	return &jwtProvider{
		prefix: prefix,
		name:   name,
	}
}
func (s *jwtProvider) GetPrefix() string {
	return s.prefix
}

func (s *jwtProvider) Get() interface{} {
	return s
}

func (s *jwtProvider) Name() string {
	return s.name
}

func (s *jwtProvider) InitFlags() {
	flag.StringVar(&s.secretKey, s.prefix+"-"+"secret-key", "ILoveGiaHan", "secret key for jwt")
}

func (s *jwtProvider) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.prefix)
	return nil
}

func (s *jwtProvider) Run() error {
	return s.Configure()
}

func (s *jwtProvider) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

type CustomClaim struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func (s *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	NewClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		Payload: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Second * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		},
	})
	token, err := NewClaim.SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   token,
		Created: time.Now(),
		Expiry:  expiry,
	}, nil
}

func (s *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*CustomClaim)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}
