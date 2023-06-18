package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GenerateToken(user_id string) (string, error) {
	viper.SetConfigFile("../config/master.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}

	token_lifespan, err := strconv.Atoi(viper.GetString("jwt.TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString("jwt.API_SECRET")))

}

func TokenValid(c *gin.Context) error {
	viper.SetConfigFile("../config/master.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.API_SECRET")), nil
	})
	fmt.Println("2:", token)
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}

	// Check token expiration
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expiration := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(expiration) {
			return errors.New("token has expired")
		}
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	return bearerToken
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	viper.SetConfigFile("../config/master.yaml")
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
