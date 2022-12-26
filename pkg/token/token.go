package token

import (
	"errors"
	"fmt"
	"foodies-crm/config"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func GenerateToken(authId int) (string, error) {
	tokenKey := config.GetString(config.TOKEN_KEY)
	tokenExp, err := strconv.Atoi(config.GetString(config.TOKEN_EXP))
	if err != nil {
		return "", err
	}

	expiredDuration := time.Duration(tokenExp) * time.Hour

	payload := Token{
		AuthId:  authId,
		Expired: time.Now().Add(expiredDuration),
	}
	claims := jwt.MapClaims{
		"payload": payload,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenString string) (*Token, error) {
	tokenKey := config.GetString(config.TOKEN_KEY)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(tokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloadInterface := claims["payload"]

		payload := Token{}

		payloadByte, err := json.Marshal(payloadInterface)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(payloadByte, &payload)
		if err != nil {
			return nil, err
		}

		now := time.Now()
		if now.After(payload.Expired) {
			return nil, errors.New("token expired")
		}
		return &payload, nil
	} else {
		return nil, errors.New("token invalid")
	}
}
